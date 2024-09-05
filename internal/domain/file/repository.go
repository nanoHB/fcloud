package file

import "context"

type FileRepository interface {
	GetFile(ctx context.Context, fileId string) (File, error)
	SaveFile(ctx context.Context, file File) error
	GetList(ctx context.Context, searchCriteria any) ([]File, error)
}
