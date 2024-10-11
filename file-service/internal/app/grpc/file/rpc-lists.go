package grpcfile

import (
	"context"

	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) ListFiles(ctx context.Context, req *filegrpc.ListFilesRequest) (*filegrpc.ListFilesResponse, error) {
    s.listLimiter <- struct{}{}
    defer func() { <-s.listLimiter }()

    files, err := s.fileOps.ListFiles(ctx)
    if err != nil {
        s.log.Error("failed to list files")
        return nil, err
    }

    return &filegrpc.ListFilesResponse{
        Files: files,
    }, nil
}
