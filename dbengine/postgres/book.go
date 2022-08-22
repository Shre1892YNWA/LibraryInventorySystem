package postgres

import (
	"errors"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

// All CRUD Methods on Books
func (d *postgresDatabase) GetAllBooks() ([]*datamodel.Book, error) {
	books := []*datamodel.Book{}

	err := d.Select(&books, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (d *postgresDatabase) GetBookById(id string) (*datamodel.Book, error) {
	book := []*datamodel.Book{}

	err := d.Select(&book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return book[0], nil
}

func (d *postgresDatabase) AddNewBook(a *datamodel.Book) error {

	if a == nil {
		return errors.New("received nil book object")
	}

	_, err := d.Exec("INSERT INTO books (id, tittle, genre, artist_id, library_id) VALUES ($1, $2, $3, $4, $5)", a.Id, a.Tittle, a.Genre, a.ArtistID, a.LibraryID)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresDatabase) UpdateExistingBook(id string, a *datamodel.Book) error {

	if a == nil {
		return errors.New("received nil book object")
	}

	_, err := d.Exec("UPDATE books SET tittle=$2, genre=$3, artist_id=$4, library_id=$5 WHERE id = $1", a.Id, a.Tittle, a.Genre, a.ArtistID, a.LibraryID)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresDatabase) DeleteBook(id string) error {
	_, err := d.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
