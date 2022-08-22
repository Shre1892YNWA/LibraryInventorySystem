package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m mongodbClient) getArtistCollection() *mongo.Collection {
	return m.Collection("Artists")
}

// All CRUD Methods on Artists
func (m mongodbClient) GetAllArtist() ([]*datamodel.Artist, error) {
	allArtists := []*datamodel.Artist{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	cur, err := m.getArtistCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var artist datamodel.Artist

		err := cur.Decode(&artist)
		if err != nil {
			return nil, err
		}

		allArtists = append(allArtists, &artist)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return allArtists, nil
}

func (m mongodbClient) GetArtistById(id string) (*datamodel.Artist, error) {
	artistFound := datamodel.Artist{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := m.getArtistCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&artistFound)
	if err != nil {
		return nil, err
	}

	return &artistFound, err

}

func (m mongodbClient) AddNewArtist(a *datamodel.Artist) error {
	if a == nil {
		return errors.New("received nil artist object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getArtistCollection().InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m mongodbClient) UpdateExistinArtist(id string, a *datamodel.Artist) error {
	if a == nil {
		return errors.New("received nil artist object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var updatedArtist datamodel.Artist
	searchFilter := bson.M{"_id": id}
	valuesToUpdate := bson.M{
		"$set": bson.M{
			"first_name":  a.FirstName,
			"last_name":   a.LastName,
			"gender":      a.Gender,
			"artist_type": a.Type,
		},
	}

	err := m.getArtistCollection().FindOneAndUpdate(ctx, searchFilter, valuesToUpdate).Decode(&updatedArtist)
	if err != nil {
		return mongo.WriteError{}
	}

	return nil
}

func (m mongodbClient) DeleteArtist(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getArtistCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
