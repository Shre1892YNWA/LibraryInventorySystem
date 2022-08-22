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

func GetAllLibrariesHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		libraries, _ := db.GetAllLibraries()
		if libraries == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		b, err := json.Marshal(libraries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}

func GetLibraryByIdHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestedId := mux.Vars(r)["id"]
		if requestedId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		library, err := db.GetLibraryById(requestedId)
		if library == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(library)
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

func CreateLibrarytHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
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

		var libraryObj datamodel.Library

		err = json.Unmarshal(b, &libraryObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		generatedId := uuid.New().String()
		libraryObj.Id = generatedId

		err = libraryObj.ValidateLibrary()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.AddNewLibrary(&libraryObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}
func UpdateLibraryHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
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

		var libraryObj datamodel.Library
		libraryObj.Id = receivedId

		err = json.Unmarshal(b, &libraryObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = libraryObj.ValidateLibrary()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateExistinLibrary(receivedId, &libraryObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}

func DeleteLibraryHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		recID := mux.Vars(r)["id"]
		if len(recID) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		err := db.DeleteLibrary(recID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
