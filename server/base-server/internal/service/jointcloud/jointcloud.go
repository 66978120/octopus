package jointcloud

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"time"

	"server/base-server/internal/data/jointcloud"
	"server/base-server/internal/data/pipeline"
	"server/common/errors"
	"strings"

	"github.com/jinzhu/copier"
)

type JointCloudService struct {
	api.UnimplementedJointCloudServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewJointCloudService(conf *conf.Bootstrap, data *data.Data) api.JointCloudServiceServer {
	s := &JointCloudService{
		conf: conf,
		data: data,
	}
	return s
}

func (s *JointCloudService) ListJointCloudDataset(ctx context.Context, req *api.ListJointCloudDatasetRequest) (*api.ListJointCloudDatasetReply, error) {
	reply, err := s.data.JointCloud.ListDataSet(ctx, &jointcloud.DataSetQuery{
		PageIndex: int(req.PageIndex),
		PageSize:  int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	dataSets := make([]*api.ListJointCloudDatasetReply_DataSet, 0)
	for _, n := range reply.List {
		dataSet := &api.ListJointCloudDatasetReply_DataSet{}
		err := copier.Copy(dataSet, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		dataSets = append(dataSets, dataSet)
	}

	return &api.ListJointCloudDatasetReply{
		TotalSize: reply.Pager.Total,
		DataSets:  dataSets,
	}, nil
}

func (s *JointCloudService) ListJointCloudDatasetVersion(ctx context.Context, req *api.ListJointCloudDatasetVersionRequest) (*api.ListJointCloudDatasetVersionReply, error) {
	reply, err := s.data.JointCloud.ListDataSetVersion(ctx, &jointcloud.DataSetVersionQuery{
		PageIndex:   int(req.PageIndex),
		PageSize:    int(req.PageSize),
		DataSetCode: req.DataSetCode,
	})
	if err != nil {
		return nil, err
	}
	versions := make([]*api.ListJointCloudDatasetVersionReply_DataSetVersion, 0)
	for _, n := range reply.List {
		version := &api.ListJointCloudDatasetVersionReply_DataSetVersion{}
		err := copier.Copy(version, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		versions = append(versions, version)
	}

	return &api.ListJointCloudDatasetVersionReply{
		TotalSize: reply.Pager.Total,
		Versions:  versions,
	}, nil
}

func (s *JointCloudService) ListJointCloudFramework(ctx context.Context, req *api.ListJointCloudFrameworkRequest) (*api.ListJointCloudFrameworkReply, error) {
	reply, err := s.data.JointCloud.ListFramework(ctx)
	if err != nil {
		return nil, err
	}
	frameworks := make([]*api.ListJointCloudFrameworkReply_Framework, 0)
	for _, n := range reply.List {
		framework := &api.ListJointCloudFrameworkReply_Framework{}
		err := copier.Copy(framework, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		frameworks = append(frameworks, framework)
	}

	return &api.ListJointCloudFrameworkReply{Frameworks: frameworks}, nil
}
func (s *JointCloudService) ListJointCloudFrameworkVersion(ctx context.Context, req *api.ListJointCloudFrameworkVersionRequest) (*api.ListJointCloudFrameworkVersionReply, error) {
	reply, err := s.data.JointCloud.ListFrameworkVersion(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	versions := make([]*api.ListJointCloudFrameworkVersionReply_Version, 0)
	for _, n := range reply.List {
		version := &api.ListJointCloudFrameworkVersionReply_Version{}
		err := copier.Copy(version, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		versions = append(versions, version)
	}

	return &api.ListJointCloudFrameworkVersionReply{Versions: versions}, nil
}
func (s *JointCloudService) ListJointCloudInterpreter(ctx context.Context, req *api.ListJointCloudInterpreterRequest) (*api.ListJointCloudInterpreterReply, error) {
	reply, err := s.data.JointCloud.ListInterpreter(ctx)
	if err != nil {
		return nil, err
	}
	interpreters := make([]*api.ListJointCloudInterpreterReply_Interpreter, 0)
	for _, n := range reply.List {
		interpreter := &api.ListJointCloudInterpreterReply_Interpreter{}
		err := copier.Copy(interpreter, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		interpreters = append(interpreters, interpreter)
	}

	return &api.ListJointCloudInterpreterReply{Interpreters: interpreters}, nil
}
func (s *JointCloudService) ListJointCloudInterpreterVersion(ctx context.Context, req *api.ListJointCloudInterpreterVersionRequest) (*api.ListJointCloudInterpreterVersionReply, error) {
	reply, err := s.data.JointCloud.ListInterpreterVersion(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	versions := make([]*api.ListJointCloudInterpreterVersionReply_Version, 0)
	for _, n := range reply.List {
		version := &api.ListJointCloudInterpreterVersionReply_Version{}
		err := copier.Copy(version, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		versions = append(versions, version)
	}

	return &api.ListJointCloudInterpreterVersionReply{Versions: versions}, nil
}

func (s *JointCloudService) checkPermForJob(ctx context.Context, job *jointcloud.JointcloudJobParam) error {

	for _, dataset := range job.DataSetVersionVoList {
		if !strings.HasPrefix(dataset.Path, "/") {
			return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
		}
	}

	if job.OutputPath != "" && !strings.HasPrefix(job.OutputPath, "/") {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	if len(job.DataSetVersionVoList) == 0 {
		job.DataSetVersionVoList = []*jointcloud.DataSetVersionVo{}
	}
	if len(job.Params) == 0 {
		job.Params = []*jointcloud.Param{}
	}
	if len(job.ResourceParams) == 0 {
		job.ResourceParams = []*jointcloud.ResourceParam{}
	}

	return nil
}

func (s *JointCloudService) TrainJob(ctx context.Context, req *api.JointCloudTrainJobRequest) (*api.JointCloudTrainJobReply, error) {
	//check 任务是否重名，联合索引。同名且未软删除，则报错。
	_, err := s.data.TrainJobDao.GetTrainJobByName(ctx, req.TaskName, req.UserId, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	trainJob := &jointcloud.TrainJob{}
	submitPara := &jointcloud.JointcloudJobParam{}
	err = copier.Copy(trainJob, req)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(submitPara, req)
	if err != nil {
		return nil, err
	}
	trainJob.Status = pipeline.PREPARING
	//各类参数校验
	err = s.checkPermForJob(ctx, submitPara)
	if err != nil {
		return nil, err
	}
	//submit job
	reply, err := s.data.JointCloud.SubmitJob(ctx, submitPara)
	if err != nil {
		return nil, err
	}
	trainJob.TaskId = reply.TaskId
	trainJob.JobId = reply.JobId
	//create recorde
	err = s.data.JointCloudDao.CreateTrainJob(ctx, trainJob)
	if err != nil {
		return nil, err
	}

	return &api.JointCloudTrainJobReply{JobId: reply.JobId, TaskId: reply.TaskId}, nil
}

func (s *JointCloudService) ListJointCloudJob(ctx context.Context, req *api.ListJointCloudJobRequest) (*api.ListJointCloudJobReply, error) {

	var totalSize int64
	ids := req.Ids
	jobList := make([]*api.JointCloudJReplyJob, 0)
	userIdMap := map[string]string{}
	spaceIdMap := map[string]string{}

	if ids == "" {
		jobs, size, err := s.data.JointCloudDao.GetTrainJobList(ctx, &jointcloud.TrainJobListQuery{
			PageIndex:   int(req.PageIndex),
			PageSize:    int(req.PageSize),
			UserId:      req.UserId,
			WorkspaceId: req.SpaceId,
		})

		if err != nil {
			return nil, err
		}

		if len(jobs) == 0 {
			return &api.ListJointCloudJobReply{
				TotalSize: 0,
				List:      jobList,
			}, nil
		}

		ids = "["
		for i, job := range jobs {
			ids = ids + job.TaskId
			if i < len(jobs)-1 {
				ids = ids + ","
			}
			userIdMap[job.TaskId] = job.UserId
			spaceIdMap[job.TaskId] = job.WorkspaceId
		}
		ids = ids + "]"
		totalSize = size
	}

	reply, err := s.data.JointCloud.ListJob(ctx, &jointcloud.JobQuery{
		PageIndex: 1,
		PageSize:  int(req.PageSize),
		Ids:       ids,
	})
	if err != nil {
		return nil, err
	}

	for _, n := range reply.List {
		job := &api.JointCloudJReplyJob{}
		err := copier.Copy(job, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		createTime, err := time.Parse("2006-01-02 15:04:05", n.CreateTime)
		if err == nil {
			job.CreateTime = createTime.Unix()
		}
		if userId, ok := userIdMap[job.TaskId]; ok {
			job.UserId = userId
		}
		if spaceId, ok := spaceIdMap[job.TaskId]; ok {
			job.WorkspaceId = spaceId
		}
		jobList = append(jobList, job)
	}

	return &api.ListJointCloudJobReply{
		TotalSize: int32(totalSize),
		List:      jobList,
	}, nil
}

func (s *JointCloudService) StopJob(ctx context.Context, req *api.JointCloudStopJobRequest) (*api.JointCloudStopJobReply, error) {
	_, err := s.data.JointCloudDao.GetTrainJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = s.data.JointCloud.StopJob(ctx, &jointcloud.StopJobRequest{TaskId: req.Id})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	err = s.data.JointCloudDao.UpdateTrainJob(ctx, &jointcloud.TrainJob{
		TaskId:      req.Id,
		Operation:   req.Operation,
		Status:      pipeline.STOPPED,
		CompletedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return &api.JointCloudStopJobReply{StoppedAt: time.Now().Unix()}, nil
}