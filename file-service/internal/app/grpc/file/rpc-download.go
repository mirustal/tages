package grpcfile

import (
	"context"

	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) DownloadFile(ctx context.Context, req *filegrpc.DownloadRequest) (*filegrpc.DownloadResponse, error) {
    s.uploadLimiter <- struct{}{}
    defer func() { <-s.uploadLimiter }()

    fileData, err := s.fileOps.DownloadFile(ctx, req.GetFileName())
    if err != nil {
        s.log.Error("failed to download file")
        return nil, err
    }

    return &filegrpc.DownloadResponse{
        FileName:   req.GetFileName(),
        FileChunk:  fileData,
    }, nil
}
