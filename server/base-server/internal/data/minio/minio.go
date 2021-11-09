package minio

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/common/errors"
	"time"

	"server/common/log"

	miniogo "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio interface {
	// 创建桶
	CreateBucket(bucketName string) error
	// 删除桶
	DeleteBucket(bucketName string) error
	// 生成上传对象url
	PresignedUploadObject(bucketName string, objectName string, domain string) (*url.URL, error)
	// 生成下载对象url
	PresignedDownloadObject(bucketName string, objectName string, domain string) (*url.URL, error)
	// 查看对象
	ListObjects(bucketName string, objectPrefix string, recurSvie bool) ([]*ObjectInfo, error)
	// 查看对象是否存在
	ObjectExist(bucketName string, objectName string) (bool, error)
}

type ObjectInfo struct {
	Name         string
	LastModified int64
	Size         int64
	ContentType  string
}

type minio struct {
	log    *log.Helper
	conf   *conf.Data
	client *miniogo.Client
}

func NewMinio(conf *conf.Data, logger log.Logger) Minio {
	client, err := miniogo.New(conf.Minio.Base.EndPoint, &miniogo.Options{
		Creds:  credentials.NewStaticV4(conf.Minio.Base.AccessKeyID, conf.Minio.Base.SecretAccessKey, ""),
		Secure: conf.Minio.Base.UseSSL,
	})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioBucketInitFailed)
		panic(err)
	}

	minio := &minio{
		log:    log.NewHelper("Minio", logger),
		conf:   conf,
		client: client,
	}

	// 创建默认的桶
	err = minio.CreateBucket(common.BUCKET)
	if err != nil {
		panic(err)
	}

	return minio
}

// 创建桶
func (m *minio) CreateBucket(bucketName string) error {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return err
	}
	if isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketExisted)
		m.log.Warnw(ctx, err)
		return nil
	}

	err = m.client.MakeBucket(ctx, bucketName, miniogo.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: false,
	})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioMakeBucketFailed)
		return err
	}

	m.log.Infof(ctx, "successfully created mybucket, bucketName=%s", bucketName)
	return nil
}

// 删除桶
func (m *minio) DeleteBucket(bucketName string) error {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return err
	}
	if !isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketNotExist)
		m.log.Errorw(ctx, err)
		return nil
	}

	err = m.client.RemoveBucket(ctx, bucketName)
	if err != nil {
		err := errors.Errorf(err, errors.ErrorMinioDeleteBucketFailed)
		return err
	}

	m.log.Infof(ctx, "successfully delete mybucket, bucketName=%s", bucketName)
	return nil
}

// 生成上传对象url
func (m *minio) PresignedUploadObject(bucketName string, objectName string, domain string) (*url.URL, error) {
	ctx := context.Background()

	uploadExpiry := time.Duration(m.conf.Minio.Business.UploadExpiry) * time.Second // duration单位是纳秒，所有得换算下
	uri, err := m.client.PresignedPutObject(ctx, bucketName, objectName, uploadExpiry)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioPresignedPutObjectFailed)
		return nil, err
	}

	m.log.Infof(ctx, "successfully PresignedPutObject, bucketName=%s|url=%s", bucketName, uri)

	if domain == "" {
		return uri, nil
	}

	domainUrl, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	uri.Path = path.Join(m.conf.Minio.Base.ProxyPath, uri.Path)
	uri.Host = domainUrl.Host
	uri.Scheme = domainUrl.Scheme
	m.log.Infof(ctx, "successfully PresignedPutObject change domain, bucketName=%s|url=%s", bucketName, uri)
	return uri, nil
}

// 生成下载对象url
func (m *minio) PresignedDownloadObject(bucketName string, objectName string, domain string) (*url.URL, error) {
	ctx := context.Background()

	reqParams := make(url.Values)
	paramKey := "response-content-disposition"
	paramVale := fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(objectName))
	reqParams.Set(paramKey, paramVale)
	downloadExpiry := time.Duration(m.conf.Minio.Business.DownloadExpiry) * time.Second // duration单位是纳秒，所有得换算下
	uri, err := m.client.PresignedGetObject(ctx, bucketName, objectName, downloadExpiry, reqParams)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioPresignedGetObjectFailed)
		return nil, err
	}

	m.log.Infof(ctx, "successfully PresignedGetObject, bucketName=%s|url=%s", bucketName, uri)

	if domain == "" {
		return uri, nil
	}

	domainUrl, err := url.Parse(domain)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorUrlParseFailed)
		return nil, err
	}
	uri.Path = path.Join(m.conf.Minio.Base.ProxyPath, uri.Path)
	uri.Host = domainUrl.Host
	uri.Scheme = domainUrl.Scheme
	m.log.Infof(ctx, "successfully PresignedGetObject change domain, bucketName=%s|url=%s", bucketName, uri)
	return uri, nil
}

// 查看对象
func (m *minio) ListObjects(bucketName string, objectPrefix string, recurSvie bool) ([]*ObjectInfo, error) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	opts := miniogo.ListObjectsOptions{
		Prefix:    objectPrefix,
		Recursive: recurSvie,
		MaxKeys:   0, // 暂时没用，先随便赋值
		UseV1:     false,
	}

	var objects []*ObjectInfo
	objectCh := m.client.ListObjects(ctx, bucketName, opts)
	for object := range objectCh {
		if object.Err != nil {
			err := errors.Errorf(object.Err, errors.ErrorMinioListObjectFailed)
			return nil, err
		}

		objects = append(objects, &ObjectInfo{
			Name:         object.Key,
			LastModified: object.LastModified.Unix(),
			Size:         object.Size,
			ContentType:  object.ContentType,
		})
	}

	m.log.Infof(ctx, "successfully ListObjects, bucketName=%s|objectPrefix=%s", bucketName, objectPrefix)
	return objects, nil
}

// 查看对象是否存在
func (m *minio) ObjectExist(bucketName string, objectName string) (bool, error) {
	ctx := context.Background()

	isExist, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckBucketExistFailed)
		return false, err
	}
	if !isExist {
		err = errors.Errorf(err, errors.ErrorMinioBucketNotExist)
		m.log.Errorw(ctx, err)
		return false, nil
	}

	_, err = m.client.StatObject(ctx, bucketName, objectName, miniogo.StatObjectOptions{})
	if err != nil {
		err = errors.Errorf(err, errors.ErrorMinioCheckObjectExistFailed)
		m.log.Errorw(ctx, err)
		return false, nil
	}

	return true, nil
}
