package postgres

import (
	"errors"
	"fmt"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

// All CRUD Methods on Artists
func (d *postgresDatabase) GetAllArtist() ([]*datamodel.Artist, error) {
	artists := []*datamodel.Artist{}

	err := d.Select(&artists, "SELECT * FROM artists")
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (d *postgresDatabase) GetArtistById(id string) (*datamodel.Artist, error) {
	artist := []*datamodel.Artist{}

	err := d.Select(&artist, "SELECT * FROM artists WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return artist[0], nil
}

func (d *postgresDatabase) AddNewArtist(a *datamodel.Artist) error {

	if a == nil {
		return errors.New("received nil artist object")
	}

	_, err := d.Exec("INSERT INTO artists (id, first_name, last_name, gender, artist_type) VALUES ($1, $2, $3, $4, $5)", a.Id, a.FirstName, a.LastName, a.Gender, a.Type)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresDatabase) UpdateExistinArtist(id string, a *datamodel.Artist) error {

	if a == nil {
		return errors.New("received nil artist object")
	}

	_, err := d.Exec("UPDATE artists SET first_name=$2, last_name=$3, gender=$4, artist_type=$5 WHERE id=$1", a.Id, a.FirstName, a.LastName, a.Gender, a.Type)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (d *postgresDatabase) DeleteArtist(id string) error {
	_, err := d.Exec("DELETE FROM artists WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
