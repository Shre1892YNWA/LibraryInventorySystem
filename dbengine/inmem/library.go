package inmem

import (
	"errors"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

func (d *inMemDB) GetAllLibraries() ([]*datamodel.Library, error) {
	var allLibraries []*datamodel.Library

	for _, v := range d.library {
		allLibraries = append(allLibraries, v)
	}

	return allLibraries, nil
}

func (d *inMemDB) GetLibraryById(id string) (*datamodel.Library, error) {

	if len(id) == 0 {
		return nil, errors.New("error id not received")
	}

	s, ok := d.library[id]
	if !ok {
		return nil, errors.New("error library does not exist")
	}

	return s, nil
}

func (d *inMemDB) AddNewLibrary(l *datamodel.Library) error {

	if l == nil {
		return errors.New(" received nil library object")
	}

	id := l.Id
	d.library[id] = l

	return nil
}

func (d *inMemDB) UpdateExistinLibrary(id string, l *datamodel.Library) error {

	if len(id) == 0 {
		return errors.New("invalid Id")
	}

	if l == nil {
		return errors.New(" received nil library object")
	}

	d.library[id] = l

	return nil
}

func (d *inMemDB) DeleteLibrary(id string) error {

	if len(id) == 0 {
		return errors.New("invalid Id")
	}

	books, _ := d.GetAllBooks()
	for _, book := range books {
		if book.LibraryID == id {
			d.DeleteBook(book.Id)
		}
	}

	delete(d.library, id)

	return nil
}
