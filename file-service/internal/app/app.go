package app

import (
	"log/slog"

	grpcapp "tages-task/file-service/internal/app/grpc"
	"tages-task/file-service/internal/modules/file"
	"tages-task/file-service/pkg/config"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, cfg *config.GRPCConfig) *App {
	fileService := file.New(log)
	grpcApp := grpcapp.New(log, fileService, cfg)

	return &App{
		GRPCServer: grpcApp,
	}
}
