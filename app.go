package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//Initialize function
func (a *App) Initialize() {
	db, err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=dbase sslmode=disable password=hello")
	if err != nil {
		log.Fatal(err)
	}

	a.DB = dbMigrate(db)

	a.Router = mux.NewRouter()
	a.initRouters()
}

// StartServer does what name implies
func (a *App) StartServer(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

//Routers Below:::
func (a *App) initRouters() {
	a.Router.HandleFunc("/api/v1/books", a.getBooks).Methods("GET")
	a.Router.HandleFunc("/api/v1/books", a.createBook).Methods("POST")
	a.Router.HandleFunc("/api/v1/books/{id}", a.updateBook).Methods("PUT")
	a.Router.HandleFunc("/api/v1/books/{id}", a.deleteBook).Methods("DELETE")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (a *App) getBooks(w http.ResponseWriter, r *http.Request) {
	loadbooksQuary(a.DB, w, r)
}

func (a *App) createBook(w http.ResponseWriter, r *http.Request) {
	createbookQuary(a.DB, w, r)
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	updatebookQuary(a.DB, w, r)
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	deletebookQuary(a.DB, w, r)
}
