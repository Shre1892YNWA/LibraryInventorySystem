package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// All CRUD Methods on Books

func (m mongodbClient) getBooksCollection() *mongo.Collection {
	return m.Collection("Books")
}
func (m mongodbClient) GetAllBooks() ([]*datamodel.Book, error) {
	allBooks := []*datamodel.Book{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	cur, err := m.getBooksCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var book datamodel.Book

		err := cur.Decode(&book)
		if err != nil {
			return nil, err
		}

		allBooks = append(allBooks, &book)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return allBooks, nil
}

func (m mongodbClient) GetBookById(id string) (*datamodel.Book, error) {
	bookFound := datamodel.Book{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err := m.getBooksCollection().FindOne(ctx, bson.M{"_id": id}).Decode(&bookFound)
	if err != nil {
		return nil, err
	}

	return &bookFound, nil
}

func (m mongodbClient) AddNewBook(a *datamodel.Book) error {
	if a == nil {
		return errors.New("received nil books object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getBooksCollection().InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m mongodbClient) UpdateExistingBook(id string, a *datamodel.Book) error {
	if a == nil {
		return errors.New("received nil books object")
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var updatedBook datamodel.Book
	searchFilter := bson.M{"_id": id}
	valuesToUpdate := bson.M{
		"$set": bson.M{
			"tittle":     a.Tittle,
			"genre":      a.Genre,
			"artist_id":  a.ArtistID,
			"library_id": a.LibraryID,
		},
	}

	err := m.getBooksCollection().FindOneAndUpdate(ctx, searchFilter, valuesToUpdate).Decode(&updatedBook)
	if err != nil {
		return mongo.WriteErrors{}
	}

	return nil
}

func (m mongodbClient) DeleteBook(id string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := m.getBooksCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
