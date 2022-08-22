package router

import (
	"errors"
	"net/http"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"

	"github.com/gorilla/mux"
)

func InitializeRoutes(db dbengine.DBEngine) (*mux.Router, error) {
	if db == nil {
		return nil, errors.New("received nil database engine")
	}

	r := mux.NewRouter()

	//Artists route handlers
	r.HandleFunc("/api/artists", GetAllArtistsHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/artists/{id}", GetArtistByIdHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/artists", CreateArtistHandler(db)).Methods(http.MethodPost)
	r.HandleFunc("/api/artists/{id}", UpdateArtistHandler(db)).Methods(http.MethodPut)
	r.HandleFunc("/api/artists/{id}", DeleteArtistHandler(db)).Methods(http.MethodDelete)

	//Books route handlers
	r.HandleFunc("/api/books", GetAllBooksHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/books/{id}", GetBookByIdHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/books", CreateBookHandler(db)).Methods(http.MethodPost)
	r.HandleFunc("/api/books/{id}", UpdateBookHandler(db)).Methods(http.MethodPut)
	r.HandleFunc("/api/books/{id}", DeleteBookHandler(db)).Methods(http.MethodDelete)

	//Libraries route handlers
	r.HandleFunc("/api/libraries", GetAllLibrariesHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/libraries/{id}", GetLibraryByIdHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/libraries", CreateLibrarytHandler(db)).Methods(http.MethodPost)
	r.HandleFunc("/api/libraries/{id}", UpdateLibraryHandler(db)).Methods(http.MethodPut)
	r.HandleFunc("/api/libraries/{id}", DeleteLibraryHandler(db)).Methods(http.MethodDelete)

	return r, nil
}
