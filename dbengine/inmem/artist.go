package inmem

import (
	"errors"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

func (d *inMemDB) GetAllArtist() ([]*datamodel.Artist, error) {
	var allArtist []*datamodel.Artist

	for _, v := range d.artist {
		allArtist = append(allArtist, v)
	}

	return allArtist, nil
}

func (d *inMemDB) GetArtistById(id string) (*datamodel.Artist, error) {
	s, ok := d.artist[id]
	if !ok {
		return nil, errors.New("error artist does not exist")
	}

	return s, nil
}

func (d *inMemDB) AddNewArtist(a *datamodel.Artist) error {

	if a == nil {
		return errors.New(" received nil artist object")
	}

	id := a.Id
	d.artist[id] = a

	return nil
}

func (d *inMemDB) UpdateExistinArtist(id string, a *datamodel.Artist) error {

	if len(id) == 0 {
		return errors.New("invalid id")
	}

	if a == nil {
		return errors.New("received nil artist object")
	}

	d.artist[id] = a

	return nil
}

func (d *inMemDB) DeleteArtist(id string) error {

	if len(id) == 0 {
		return errors.New("invalid id")
	}

	books, _ := d.GetAllBooks()
	for _, book := range books {
		if book.ArtistID == id {
			d.DeleteBook(book.Id)
		}
	}

	delete(d.artist, id)

	return nil
}
