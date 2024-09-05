package application

import (
	"context"
	"fmt"

	file "github.com/nanoHB/fcloud/internal/domain/file"
	fileMongo "github.com/nanoHB/fcloud/internal/infrastructure/mongo"
)

type FileService interface {
	GetListFiles(userId int) (string, error)
}

type FileServiceDefault struct {
	FileRepository file.FileRepository
}

func NewFileServiceMongo(dbService *fileMongo.Service) FileServiceDefault {
	mongoFileRepository := fileMongo.NewFileRepository(dbService)
	return FileServiceDefault{FileRepository: mongoFileRepository}
}

func (f *FileServiceDefault) GetListFiles(userId int) ([]file.File, error) {
	ctx := context.TODO()
	listFile, err := f.FileRepository.GetList(ctx, fmt.Sprintf("userId: %d", userId))
	if err != nil {

	}

	return listFile, err
}
