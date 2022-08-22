package dbengine

import "github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"

type DBEngine interface {
	
	// All CRUD Methods on Books
	GetAllBooks() ([]*datamodel.Book, error)
	GetBookById(id string) (*datamodel.Book, error)
	AddNewBook(a *datamodel.Book) error
	UpdateExistingBook(id string, a *datamodel.Book) error
	DeleteBook(id string) error

	// All CRUD Methods on Artists
	GetAllArtist() ([]*datamodel.Artist, error)
	GetArtistById(id string) (*datamodel.Artist, error)
	AddNewArtist(a *datamodel.Artist) error
	UpdateExistinArtist(id string, a *datamodel.Artist) error
	DeleteArtist(id string) error

	// All CRUD Methods on Library
	GetAllLibraries() ([]*datamodel.Library, error)
	GetLibraryById(id string) (*datamodel.Library, error)
	AddNewLibrary(a *datamodel.Library) error
	UpdateExistinLibrary(id string, a *datamodel.Library) error
	DeleteLibrary(id string) error
}
