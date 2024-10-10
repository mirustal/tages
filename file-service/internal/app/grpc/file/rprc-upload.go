package grpcfile

import (
	filegrpc "tages-task/file-service/pkg/pb"
)

func (s *serverAPI) UploadFile(filegrpc.FileService_UploadFileServer) error {

	return nil
}
