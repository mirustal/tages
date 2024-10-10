package file

import (
	"context"
	"log/slog"
)

type File struct {
	log *slog.Logger
}

func New(log *slog.Logger) *File {
	return &File{
		log: log,
	}
}


func (f *File) UploadFile(context.Context) error {
	const op = "internal.modules.file.uploadFile"
	log := f.log.With(
		slog.String("op", op),
	)
	log.Info("accept")
	return nil
}

func (f *File) ListFiles(context.Context) error {
	const op = "internal.modules.file.listFiles"
	log := f.log.With(
		slog.String("op", op),
	)
	log.Info("accept")
	return nil
}

func (f *File) DownloadFile(context.Context) error {
	const op = "internal.modules.file.downloadFile"
	log := f.log.With(
		slog.String("op", op),
	)
	log.Info("accept")
	return nil
}