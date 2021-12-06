package platform

import (
	"context"
	stderrors "errors"
	"fmt"
	model "server/base-server/internal/data/dao/model/platform"
	"server/common/errors"
	"server/common/transaction"
	"server/common/utils"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type PlatformDao interface {
	CreatePlatform(ctx context.Context, platform *model.Platform) error
	ListPlatform(ctx context.Context, query *model.PlatformQuery) ([]*model.Platform, int64, error)
	// 空值字段不更新
	UpdatePlatformById(ctx context.Context, platform *model.Platform) error

	CreatePlatformStorageConfig(ctx context.Context, platformStorageConfig *model.PlatformStorageConfig) error
	ListPlatformStorageConfig(ctx context.Context, query *model.PlatformStorageConfigQuery) ([]*model.PlatformStorageConfig, int64, error)
	DeletePlatformStorageConfig(ctx context.Context, platformId string, name string) error

	UpdatePlatformConfig(ctx context.Context, platformId string, config map[string]string) error
	GetPlatformConfig(ctx context.Context, platformId string) (map[string]string, error)
}

type platformDao struct {
	db transaction.GetDB
}

func NewPlatformDao(db *gorm.DB) PlatformDao {
	return &platformDao{
		db: func(ctx context.Context) *gorm.DB {
			return transaction.GetDBFromCtx(ctx, db)
		},
	}
}

func (d *platformDao) CreatePlatform(ctx context.Context, platform *model.Platform) error {
	db := d.db(ctx)
	res := db.Create(platform)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *platformDao) ListPlatform(ctx context.Context, query *model.PlatformQuery) ([]*model.Platform, int64, error) {
	db := d.db(ctx)
	platforms := make([]*model.Platform, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.Name != "" {
		querySql += " and name = ? "
		params = append(params, query.Name)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.Platform{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&platforms)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return platforms, totalSize, nil
}

func (d *platformDao) UpdatePlatformById(ctx context.Context, platform *model.Platform) error {
	db := d.db(ctx)
	if platform.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Where("id = ?", platform.Id).Updates(platform)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *platformDao) CreatePlatformStorageConfig(ctx context.Context, platformStorageConfig *model.PlatformStorageConfig) error {
	db := d.db(ctx)
	res := db.Create(platformStorageConfig)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *platformDao) ListPlatformStorageConfig(ctx context.Context, query *model.PlatformStorageConfigQuery) ([]*model.PlatformStorageConfig, int64, error) {
	db := d.db(ctx)
	platformStorageConfigs := make([]*model.PlatformStorageConfig, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.Name != "" {
		querySql += " and name = ? "
		params = append(params, query.Name)
	}

	if query.PlatformId != "" {
		querySql += " and platform_id = ? "
		params = append(params, query.PlatformId)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.PlatformStorageConfig{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&platformStorageConfigs)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return platformStorageConfigs, totalSize, nil
}

func (d *platformDao) DeletePlatformStorageConfig(ctx context.Context, platformId string, name string) error {
	db := d.db(ctx)
	if platformId == "" || name == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("platform_id = ? and name = ?", platformId, name).Limit(1).Delete(&model.PlatformStorageConfig{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *platformDao) UpdatePlatformConfig(ctx context.Context, platformId string, config map[string]string) error {
	db := d.db(ctx)
	if platformId == "" || len(config) == 0 {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	configUp := make(map[string]string)
	for k, v := range config {
		if v != "" {
			configUp[k] = v
		}
	}

	c := &model.PlatformConfig{PlatformId: platformId, Config: configUp}
	res := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(c)
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBUpdateFailed)
	}

	return nil
}

func (d *platformDao) GetPlatformConfig(ctx context.Context, platformId string) (map[string]string, error) {
	db := d.db(ctx)
	c := &model.PlatformConfig{}

	res := db.First(c, "platform_id = ?", platformId)
	if res.Error != nil && !stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return c.Config, nil
}