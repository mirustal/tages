package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"tages-task/file-service/internal/app"
	"tages-task/file-service/pkg/config"
	"tages-task/file-service/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatalf("fail load config: %v", err)
	}

	logger := logger.LogInit(cfg.ModeLog)

	app := app.New(logger.Log, cfg.GRPC)

	go app.GRPCServer.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Log.Info("recieved signal", <-c)
	app.GRPCServer.Stop()
	logger.Log.Info("file service stop")
}
