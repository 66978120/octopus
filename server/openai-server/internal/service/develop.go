package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/session"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type DevelopService struct {
	api.UnimplementedDevelopServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewDevelopService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.DevelopServer {
	return &DevelopService{
		conf: conf,
		log:  log.NewHelper("DevelopService", logger),
		data: data,
	}
}

func (s *DevelopService) CreateNotebook(ctx context.Context, req *api.CreateNotebookRequest) (*api.CreateNotebookReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.CreateNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.DevelopClient.CreateNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CreateNotebookReply{Id: innerReply.Id}, nil
}

func (s *DevelopService) checkPermission(ctx context.Context, notebookId string, userId string) error {
	reply, err := s.data.DevelopClient.GetNotebook(ctx, &innerapi.GetNotebookRequest{Id: notebookId})
	if err != nil {
		return err
	}

	if reply.Notebook.UserId != userId {
		return errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	return nil
}

func (s *DevelopService) StartNotebook(ctx context.Context, req *api.StartNotebookRequest) (*api.StartNotebookReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkPermission(ctx, req.Id, session.UserId)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.StartNotebook(ctx, &innerapi.StartNotebookRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &api.StartNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) StopNotebook(ctx context.Context, req *api.StopNotebookRequest) (*api.StopNotebookReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkPermission(ctx, req.Id, session.UserId)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.StopNotebook(ctx, &innerapi.StopNotebookRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &api.StopNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) DeleteNotebook(ctx context.Context, req *api.DeleteNotebookRequest) (*api.DeleteNotebookReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkPermission(ctx, req.Id, session.UserId)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.DeleteNotebook(ctx, &innerapi.DeleteNotebookRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &api.DeleteNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) ListNotebook(ctx context.Context, req *api.ListNotebookRequest) (*api.ListNotebookReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.DevelopClient.ListNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListNotebookReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
