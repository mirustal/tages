package grpcfile

import (
	"context"

	"google.golang.org/grpc"

	filegrpc "tages-task/file-service/pkg/pb"
)

const (
	bearerPrefix = "Bearer "
)

type serverAPI struct {
	filegrpc.UnimplementedFileServiceServer
	fileOps FileOperations
}

func Register(gRPC *grpc.Server, file FileOperations) {
	filegrpc.RegisterFileServiceServer(gRPC, &serverAPI{
		fileOps: file,
	})
}

type FileOperations interface {
    UploadFile(ctx context.Context) error
    ListFiles(ctx context.Context) (error)
    DownloadFile(ctx context.Context) error
}
