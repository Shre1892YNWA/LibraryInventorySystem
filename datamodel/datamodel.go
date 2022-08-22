package datamodel

import (
	"fmt"
	"strings"
)

type Artist struct {
	Id        string `json:"id" db:"id" bson:"_id"`
	FirstName string `json:"first_name" db:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" db:"last_name" bson:"last_name"`
	Gender    string `json:"gender" db:"gender" bson:"gender"`
	Type      string `json:"artist_type" db:"artist_type" bson:"artist_type"`
}

type Library struct {
	Id   string `json:"id" db:"id" bson:"_id"`
	Name string `json:"library_name" db:"library_name" bson:"library_name"`
	City string `json:"city" db:"city" bson:"city"`
}

type Book struct {
	Id        string `json:"id" db:"id" bson:"_id"`
	Tittle    string `json:"tittle" db:"tittle" bson:"tittle"`
	Genre     string `json:"genre" db:"genre" bson:"genre"`
	ArtistID  string `json:"artist_id" db:"artist_id" bson:"artist_id"`
	LibraryID string `json:"library_id" db:"library_id" bson:"library_id"`
}

func (a *Artist) ValidateArtist() error {
	var invalidFields []string
	var err error

	if len(a.Id) == 0 {
		invalidFields = append(invalidFields, "id %s", a.Id)
	}

	if len(a.FirstName) == 0 {
		invalidFields = append(invalidFields, "first name %s", a.FirstName)
	}

	if len(a.LastName) == 0 {
		invalidFields = append(invalidFields, "last name %s", a.LastName)
	}

	if a.Gender != "Male" && a.Gender != "Female" {
		invalidFields = append(invalidFields, "gender %s", a.Gender)
	}

	if len(a.Type) == 0 {
		invalidFields = append(invalidFields, "author type %s", a.Type)
	}

	if len(invalidFields) > 0 {
		err = fmt.Errorf("invalid data found at %s", strings.Join(invalidFields, ", "))
	}

	return err
}

func (a *Book) ValidateBook() error {
	var invalidFields []string
	var err error

	if len(a.Id) == 0 {
		invalidFields = append(invalidFields, "invalid Id %s", a.Id)
	}

	if len(a.Tittle) == 0 {
		invalidFields = append(invalidFields, "invalid first name %s", a.Tittle)
	}

	if len(a.Genre) == 0 {
		invalidFields = append(invalidFields, "invalid last name %s", a.Genre)
	}

	if len(a.ArtistID) == 0 {
		invalidFields = append(invalidFields, "invalid author tpyr %s", a.ArtistID)
	}

	if len(a.LibraryID) == 0 {
		invalidFields = append(invalidFields, "invalid author tpyr %s", a.LibraryID)
	}

	if len(invalidFields) > 0 {
		err = fmt.Errorf("invalid data found at %s", strings.Join(invalidFields, ", "))
	}

	return err
}

func (a *Library) ValidateLibrary() error {
	var invalidFields []string
	var err error

	if len(a.Id) == 0 {
		invalidFields = append(invalidFields, "invalid Id %s", a.Id)
	}

	if len(a.City) == 0 {
		invalidFields = append(invalidFields, "invalid city %s", a.City)
	}

	if len(a.Name) == 0 {
		invalidFields = append(invalidFields, "invalid last name %s", a.Name)
	}

	if len(invalidFields) > 0 {
		err = fmt.Errorf("invalid data found at %s", strings.Join(invalidFields, ", "))
	}

	return err
}
