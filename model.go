package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//Book Struct/Model
type book struct {
	ID          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

//AutoMigrate creates table of book with struct values in database
func dbMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&book{})
	return db
}

//DataBase Queries:::
func loadbooksQuary(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books []book
	if err := db.Find(&books).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)

}

func createbookQuary(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var books book

	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	books.PublishedAt = timeFunc(books.PublishedAt)

	if err := db.Save(&books).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, books)
}

func timeFunc(a string) string {
	layoutISO := "2006-01-02"
	nt, _ := time.Parse(layoutISO, a)
	return nt.String()
}

func idCheck(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *book {
	books := book{}
	if err := db.First(&books, book{ID: id}).Error; err != nil {
		return nil
	}
	return &books
}

func updatebookQuary(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aa := vars["id"]
	id, _ := strconv.Atoi(aa)
	books := book{}
	idCheck := idCheck(db, id, w, r)
	if idCheck == nil {
		respondWithError(w, http.StatusNotFound, "ID not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		respondWithError(w, http.StatusNoContent, err.Error())
		return
	}
	defer r.Body.Close()

	books.ID = id

	if err := db.Save(&books).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, books)
}

func deletebookQuary(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books book
	if err := db.First(&books, params["id"]).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	db.Delete(&books)

	var book []book
	if err := db.Find(&book).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, "DELETED")
}
