package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbName = "shre"

type mongodbClient struct {
	*mongo.Database
}

func GetMongodbEngine(mongodbUrl string) (dbengine.DBEngine, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	clientOptions := options.Client().ApplyURI(mongodbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	mDB := client.Database(mongodbName)
	if mDB == nil {
		return nil, errors.New("database not found")
	}

	return &mongodbClient{mDB}, err
}
