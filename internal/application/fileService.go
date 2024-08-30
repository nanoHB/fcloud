package application

import (
	file "github.com/nanoHB/fcloud/internal/domain/file"
)

type FileService interface {
	GetListFiles(userId int) (string, error)
}

type FileServiceDefault struct {
	FileRepository file.FileRepository
}
