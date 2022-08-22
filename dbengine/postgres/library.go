package postgres

import (
	"errors"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
)

// All CRUD Methods on Library
func (d *postgresDatabase) GetAllLibraries() ([]*datamodel.Library, error) {
	libraries := []*datamodel.Library{}

	err := d.Select(&libraries, "SELECT * FROM libraries")
	if err != nil {
		return nil, err
	}

	return libraries, nil
}

func (d *postgresDatabase) GetLibraryById(id string) (*datamodel.Library, error) {
	library := []*datamodel.Library{}

	err := d.Select(&library, "SELECT * FROM libraries WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return library[0], nil
}

func (d *postgresDatabase) AddNewLibrary(a *datamodel.Library) error {

	if a == nil {
		return errors.New("received nil libraries object")
	}

	_, err := d.Exec("INSERT INTO libraries (id, city, library_name ) VALUES ($1, $2, $3)", a.Id, a.City, a.Name)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresDatabase) UpdateExistinLibrary(id string, a *datamodel.Library) error {

	if a == nil {
		return errors.New("received nil libraries object")
	}

	_, err := d.Exec("UPDATE libraries SET city=$2, library_name=$3 WHERE id=$1", a.Id, a.City, a.Name)
	if err != nil {
		return err
	}

	return nil
}

func (d *postgresDatabase) DeleteLibrary(id string) error {

	_, err := d.Exec("DELETE FROM libraries WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
