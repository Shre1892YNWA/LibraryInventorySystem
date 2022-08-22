package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/Shre1892YNWA/LibraryInventorySystem/datamodel"
	"github.com/Shre1892YNWA/LibraryInventorySystem/dbengine"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetAllArtistsHandler(db dbengine.DBEngine) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		artists, _ := db.GetAllArtist()
		if artists == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		b, err := json.Marshal(artists)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}

func GetArtistByIdHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		requestedId := mux.Vars(r)["id"]
		if requestedId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		artist, err := db.GetArtistById(requestedId)
		if artist == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(artist)
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

func CreateArtistHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
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

		var artistObj datamodel.Artist
		generatedId := uuid.New().String()
		artistObj.Id = generatedId

		err = json.Unmarshal(b, &artistObj)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		err = artistObj.ValidateArtist()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.AddNewArtist(&artistObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}

}

func UpdateArtistHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
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

		var artistObj datamodel.Artist
		artistObj.Id = receivedId

		err = json.Unmarshal(b, &artistObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = artistObj.ValidateArtist()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateExistinArtist(receivedId, &artistObj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

}

func DeleteArtistHandler(db dbengine.DBEngine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		recID := mux.Vars(r)["id"]
		if len(recID) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		err := db.DeleteArtist(recID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

}
