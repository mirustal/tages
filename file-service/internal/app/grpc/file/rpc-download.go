package grpcfile

import (
	"context"

	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) DownloadFile(ctx context.Context, req *filegrpc.DownloadRequest) (*filegrpc.DownloadResponse, error) {

	return &filegrpc.DownloadResponse{}, nil
}