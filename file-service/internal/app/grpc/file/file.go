package grpcfile

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"

	"tages-task/file-service/pkg/config"
	filegrpc "tages-task/file-service/pkg/pb"
)


type serverAPI struct {
	filegrpc.UnimplementedFileServiceServer
	fileOps 		FileOperations
	log           *slog.Logger
	cfg		*config.GRPCConfig
	uploadLimiter chan struct{}
    listLimiter   chan struct{}
}

func Register(gRPC *grpc.Server, log *slog.Logger, cfg *config.GRPCConfig, file FileOperations) {
	filegrpc.RegisterFileServiceServer(gRPC, &serverAPI{
		fileOps: file,
		log: log,
		cfg: cfg,
		uploadLimiter: make(chan struct{}, 10),
        listLimiter:   make(chan struct{}, 100), 
	})
}

type FileOperations interface {
    UploadFile(ctx context.Context, stream filegrpc.FileService_UploadFileServer) (string, error)
    ListFiles(ctx context.Context) ([]*filegrpc.FileMetadata, error)
    DownloadFile(ctx context.Context, fileName string) ([]byte, error)
}


