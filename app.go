package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-REST-simple-product/models"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

// App exposes references to the Router and Database the application uses
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize sets up the connection to the database
func (a *App) Initialize(dialect, connectionInfo string) {
	var err error
	a.DB, err = sql.Open(dialect, connectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the server
func (a *App) Run(addr int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", addr), a.Router))
}

//ByID GET /product/:id
func (a *App) ByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.ByID(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)

}

//Index GET /products
// /products?start=10 etc
func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := models.ByRange(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Create POST /product creates a new product
func (a *App) Create(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.Create(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

// Update PUT /product/:id updates the given product
func (a *App) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.Product{ID: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.Update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

//Delete DELETE /product/:id deletes the given product
func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.Delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Is that all, stranger?")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/products", a.Index).Methods("GET")
	a.Router.HandleFunc("/products", a.Create).Methods("POST")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.ByID).Methods("GET")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.Update).Methods("PUT")
	a.Router.HandleFunc("/products/{id:[0-9]+}", a.Delete).Methods("DELETE")
}
