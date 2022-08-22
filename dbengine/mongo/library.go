package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m mongodbClient) getLibrariesCollections() *mongo.Collection {
	return m.Collection("Libraries")
}

// All CRUD Methods on Library
func (m mongodbClient) GetAllLibraries() ([]*datamodel.Library, error) {
	allLibraries := []*datamodel.Library{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	cur, err := m.getLibrariesCollections().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var library datamodel.Library

		err := cur.Decode(&library)
		if err != nil {
			return nil, err
		}

		allLibraries = append(allLibraries, &library)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return allLibraries, nil
}

func (m mongodbClient) GetLibraryById(id string) (*datamodel.Library, error) {
	libraryFound := datamodel.Library{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := m.getLibrariesCollections().FindOne(ctx, bson.M{"_id": id}).Decode(&libraryFound)

	if err != nil {
		return nil, err
	}

	return &libraryFound, nil
}

func (m mongodbClient) AddNewLibrary(a *datamodel.Library) error {
	if a == nil {
		return errors.New("received nil library object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getLibrariesCollections().InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m mongodbClient) UpdateExistinLibrary(id string, a *datamodel.Library) error {
	if a == nil {
		return errors.New("received nil library object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var updatedLibrary datamodel.Library
	searchFilter := bson.M{"_id": id}
	valuesToUpdate := bson.M{
		"$set": bson.M{
			"library_name": a.Name,
			"city":         a.City,
		},
	}

	err := m.getLibrariesCollections().FindOneAndUpdate(ctx, searchFilter, valuesToUpdate).Decode(&updatedLibrary)
	if err != nil {
		return mongo.WriteError{}
	}

	return nil
}

func (m mongodbClient) DeleteLibrary(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getLibrariesCollections().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
