package inmem

import (
	"errors"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

func (d *inMemDB) GetAllBooks() ([]*datamodel.Book, error) {
	var allLibraries []*datamodel.Book

	for _, v := range d.book {
		allLibraries = append(allLibraries, v)
	}

	return allLibraries, nil
}

func (d *inMemDB) GetBookById(id string) (*datamodel.Book, error) {

	if len(id) == 0 {
		return nil, errors.New("error id not received")
	}

	return d.book[id], nil
}

func (d *inMemDB) AddNewBook(l *datamodel.Book) error {

	if l == nil {
		return errors.New(" received nil Book object")
	}

	id := l.Id
	d.book[id] = l

	return nil
}

func (d *inMemDB) UpdateExistingBook(id string, l *datamodel.Book) error {

	if len(id) == 0 {
		return errors.New("invalid Id")
	}

	if l == nil {
		return errors.New(" received nil Book object")
	}

	d.book[id] = l

	return nil
}

func (d *inMemDB) DeleteBook(id string) error {

	if len(id) == 0 {
		return errors.New("invalid Id")
	}

	delete(d.book, id)

	return nil
}
