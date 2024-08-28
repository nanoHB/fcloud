package mongo

import (
	"context"

	file "github.com/nanoHB/fcloud/internal/domain/file"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FileRepository struct {
	mongoService service
}

func NewFileRepository(service *service) FileRepository {
	return FileRepository{mongoService: *service}
}

func (f FileRepository) GetFile(ctx context.Context, fileId string) (file.File, error) {
	mongoClient := f.mongoService.db
	filesCollection := mongoClient.Database("fcloud").Collection("file")
	objectId, err := primitive.ObjectIDFromHex(fileId)
	filter := bson.M{"_id": objectId}
	var file file.File

	err = filesCollection.FindOne(ctx, filter).Decode(&file)

	return file, err
}

func (f FileRepository) SaveFile(ctx context.Context, file file.File) error {
	mongoClient := f.mongoService.db
	filesCollection := mongoClient.Database("fcloud").Collection("file")
	updateOption := options.Update().SetUpsert(true)
	objectId, err := primitive.ObjectIDFromHex(file.Id)

	filter := bson.M{"_id": objectId}

	_, err = filesCollection.UpdateOne(ctx, filter, file, updateOption)

	return err
}

func (f FileRepository) GetList(ctx context.Context, searchCriteria any) ([]file.File, error) {
	mongoClient := f.mongoService.db
	filesCollection := mongoClient.Database("fcloud").Collection("file")
	cursor, err := filesCollection.Find(ctx, searchCriteria)
	defer cursor.Close(ctx)

	var files []file.File

	for cursor.Next(ctx) {
		var file file.File
		err := cursor.Decode(&file)
		if err != nil {
			continue
		}
		files = append(files, file)
	}

	return files, err
}
