package file

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	filegrpc "tages-task/file-service/pkg/pb"
)

type File struct {
	log *slog.Logger
	storagePath string
}

func New(log *slog.Logger, storagePath string) *File {
	return &File{
		log: log,
		storagePath: storagePath,
	}
}
// not work delete this 
func (f *File) UploadFile(ctx context.Context, stream filegrpc.FileService_UploadFileServer) (string, error) {
    var fileName string
    var fileData []byte

    if err := os.MkdirAll(f.storagePath, os.ModePerm); err != nil {
        return "", fmt.Errorf("failed to create storage directory: %w", err)
    }

    for {
        req, err := stream.Recv()
		fileName = req.GetFileName()
		fmt.Println("filename: ", fileName)
        if err == io.EOF {
            break
        }
        if err != nil {
            return "", fmt.Errorf("failed to receive file chunk: %w", err)
        }

        fileData = append(fileData, req.GetFileChunk()...)
    }
    filePath := filepath.Join(f.storagePath, fileName)
    if err := os.WriteFile(filePath, fileData, 0644); err != nil {
        return "", fmt.Errorf("failed to save file: %w", err)
    }

    return fileName, nil
}

func (f *File) ListFiles(ctx context.Context) ([]*filegrpc.FileMetadata, error) {
    var files []*filegrpc.FileMetadata

    err := filepath.Walk(f.storagePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            files = append(files, &filegrpc.FileMetadata{
                FileName:   info.Name(),
                CreatedAt:  info.ModTime().Format(time.RFC3339),
                UpdatedAt:  info.ModTime().Format(time.RFC3339), 
            })
        }
        return nil
    })

    if err != nil {
        return nil, fmt.Errorf("failed to list files: %w", err)
    }

    return files, nil
}

func (f *File) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
    filePath := filepath.Join(f.storagePath, fileName)
    

    fileData, err := os.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }

    return fileData, nil
}
