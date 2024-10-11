package grpcfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) UploadFile(stream filegrpc.FileService_UploadFileServer) error {
	s.uploadLimiter <- struct{}{}
	defer func() { <-s.uploadLimiter }()

	var fileName string
	var fileData []byte

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Error("failed to receive file chunk")
			return err
		}

		if fileName == "" {
			fileName = req.GetFileName()
		}
		fileData = append(fileData, req.GetFileChunk()...)
	}
	filePath := filepath.Join(s.cfg.StoragePath, fileName)
    if err := os.WriteFile(filePath, fileData, 0644); err != nil {
        return fmt.Errorf("failed to save file: %w", err)
    }

	return stream.SendAndClose(&filegrpc.UploadResponse{
		FileName: fileName,
		Message:  "File uploaded successfully",
	})
}
