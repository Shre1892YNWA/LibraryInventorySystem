package inmem

import (
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"
)

type inMemDB struct {
	artist  map[string]*datamodel.Artist
	library map[string]*datamodel.Library
	book    map[string]*datamodel.Book
}

func GetInMemEngine() dbengine.DBEngine {
	return &inMemDB{
		artist:  map[string]*datamodel.Artist{},
		library: map[string]*datamodel.Library{},
		book:    map[string]*datamodel.Book{},
	}
}
