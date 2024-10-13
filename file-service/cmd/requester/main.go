package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"

	filegrpc "tages-task/file-service/pkg/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := filegrpc.NewFileServiceClient(conn)

	fmt.Println("Выберите действие:")
	fmt.Println("1 - Скачать файл")
	fmt.Println("2 - Получить список файлов")
	fmt.Println("3 - Загрузить файл")

	for {
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			fmt.Print("Введите имя файла для скачивания: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			downloadFile(client, fileName)

		case "2":
			listFiles(client)

		case "3":
			fmt.Print("Введите путь к файлу для загрузки(полный путь с файлом): ")
			filePath, _ := reader.ReadString('\n')
			filePath = strings.TrimSpace(filePath)
			fmt.Print("Введите название файла для загрузки: ")
			fileName, _ := reader.ReadString('\n')
			fileName = strings.TrimSpace(fileName)
			uploadFile(client, filePath, fileName)

		case "q":
			break

		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1, 2 или 3.")
		}
	}
}

func downloadFile(client filegrpc.FileServiceClient, fileName string) {
	req := &filegrpc.DownloadRequest{
		FileName: fileName,
	}

	res, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		log.Fatalf("Error downloading file: %v", err)
	}

	fileData := res.FileChunk

	downloadDir := "download"
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		err := os.Mkdir(downloadDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create download directory: %v", err)
		}
	}

	filePath := filepath.Join(downloadDir, fileName)
	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}
	log.Printf("File successfully downloaded and saved as %s", filePath)
}

func listFiles(client filegrpc.FileServiceClient) {
	req := &filegrpc.ListFilesRequest{}

	res, err := client.ListFiles(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting file list: %v", err)
	}

	for _, file := range res.Files {
		fmt.Printf("File: %s, Created: %s, Updated: %s\n", file.FileName, file.CreatedAt, file.UpdatedAt)
	}
}

func uploadFile(client filegrpc.FileServiceClient, filePath string, fileName string) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}
	if fileInfo.IsDir() {
		log.Fatalf("The path %s is a directory, please provide a file path", filePath)
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	if fileName == "" {
		fileName = filepath.Base(filePath)
	}
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("Failed to open upload stream: %v", err)
	}

	const chunkSize = 1024 * 1024
	for currentByte := 0; currentByte < len(fileData); currentByte += chunkSize {
		endByte := currentByte + chunkSize
		if endByte > len(fileData) {
			endByte = len(fileData)
		}

		req := &filegrpc.UploadRequest{
			FileName:  fileName,
			FileChunk: fileData[currentByte:endByte],
		}

		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Failed to send file chunk: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to complete file upload: %v", err)
	}

	log.Printf("File uploaded: %s, Message: %s", res.FileName, res.Message)
}
