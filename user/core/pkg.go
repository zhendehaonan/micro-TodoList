package core

import (
	"context"
	service "github.com/micro/go-micro/v2/config/source/service/proto"
)

type UserService struct {
}

func (s *UserService) Create(ctx context.Context, request *service.CreateRequest, response *service.CreateResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Update(ctx context.Context, request *service.UpdateRequest, response *service.UpdateResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Delete(ctx context.Context, request *service.DeleteRequest, response *service.DeleteResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) List(ctx context.Context, request *service.ListRequest, response *service.ListResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Read(ctx context.Context, request *service.ReadRequest, response *service.ReadResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) Watch(ctx context.Context, request *service.WatchRequest, stream service.Config_WatchStream) error {
	//TODO implement me
	panic("implement me")
}
