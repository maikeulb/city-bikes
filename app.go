package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/networks", a.getNetworks).Methods("GET")
	// a.Router.HandleFunc("/networks/{id}", a.getNetwork).Methods("GET")
}

func (a *App) getNetworks(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://api.citybik.es/v2/networks")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject)
	fmt.Println(len(responseObject.Networks))

	respondWithJSON(w, http.StatusOK, responseObject)
}

// func (a *App) getNetwork(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
// 		return
// 	}

// 	n := network{ID: id}
// 	if err := p.getNetwork(); err != nil {
// 	    respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, p)
// }

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
