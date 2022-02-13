package model

import (
	"competition/dockerApi"
	"context"
)

type DockerServer struct {
}

func (d *DockerServer) CreatContainMessage(ctx context.Context, request *CreatContainRequest) (*CreatContainResponse, error) {
	ConId := dockerApi.CreatContainer(request.Alias, request.Images, request.Port1, request.Port2, request.SqlPort)
	return &CreatContainResponse{ContainerID: ConId}, nil
}

func (d *DockerServer) GetStartContainUrl(ctx context.Context, Contain *StartContainRequest) (*StartContainResponse, error) {
	err := dockerApi.StartContainer(Contain.ContainerID)
	if err != nil {
		return &StartContainResponse{ContainUrl: "err"}, err
	} else {
		return &StartContainResponse{ContainUrl: Contain.ContainerID}, nil
	}
}
