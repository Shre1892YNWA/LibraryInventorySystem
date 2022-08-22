package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllBooksHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		books, _ := db.GetAllBooks()
		if books == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		b, err := json.Marshal(books)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}

func GetBookByIdHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestedId := mux.Vars(r)["id"]
		if requestedId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		book, err := db.GetBookById(requestedId)
		if book == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(book)
		if b == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}

}

func CreateBookHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := ioutil.ReadAll(r.Body)
		if b == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var bookObj datamodel.Book

		err = json.Unmarshal(b, &bookObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		generatedId := uuid.New().String()
		bookObj.Id = generatedId

		err = bookObj.ValidateBook()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.AddNewBook(&bookObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
func UpdateBookHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		receivedId := mux.Vars(r)["id"]
		if len(receivedId) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		if b == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var bookObj datamodel.Book
		bookObj.Id = receivedId

		err = json.Unmarshal(b, &bookObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = bookObj.ValidateBook()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateExistingBook(receivedId, &bookObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}

func DeleteBookHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		recID := mux.Vars(r)["id"]
		if len(recID) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		err := db.DeleteBook(recID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
