package grpcfile

import (
	"context"

	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) ListsFiles(ctx context.Context, req *filegrpc.ListFilesRequest) (*filegrpc.ListFilesResponse, error) {

	return &filegrpc.ListFilesResponse{}, nil
}
