package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maikeulb/city-bike/redis"

	"github.com/go-redis/cache"
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
	a.Router.HandleFunc("/api/networks/{id}", a.getNetwork).Methods("GET")
}

func (a *App) getNetworks(w http.ResponseWriter, r *http.Request) {

	queryStr := "all_networks"

	var responseObject NetworksResponse
	startQuery := time.Now()
	if err := redis.Codec.Get(queryStr, &responseObject); err != nil {

		log.Println("Remote networks")
		response, err := http.Get("http://api.citybik.es/v2/networks")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &responseObject)

		updateCacheNetworksResponse(queryStr, responseObject)
	} else {
		log.Println("Cached networks")
	}
	endQuery := time.Now()
	log.Println("Got networks in ", endQuery.Sub(startQuery), " seconds")

	respondWithJSON(w, http.StatusOK, responseObject)
}

func (a *App) getNetwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	queryStr := id

	var responseObject NetworkResponse
	startQuery := time.Now()
	if err := redis.Codec.Get(queryStr, &responseObject); err != nil {

		log.Println("Remote network location")
		response, err := http.Get("http://api.citybik.es/v2/networks/" + id)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(responseData, &responseObject)

		updateCacheNetworkResponse(queryStr, responseObject)
	} else {
		log.Println("Cached network location")
	}
	endQuery := time.Now()
	log.Println("Got network location in ", endQuery.Sub(startQuery), " seconds")

	respondWithJSON(w, http.StatusOK, responseObject)
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

func updateCacheNetworksResponse(key string, responseObject NetworksResponse) {
	redis.Codec.Set(&cache.Item{
		Key:        key,
		Object:     responseObject,
		Expiration: time.Hour,
	})
}

func updateCacheNetworkResponse(key string, responseObject NetworkResponse) {
	redis.Codec.Set(&cache.Item{
		Key:        key,
		Object:     responseObject,
		Expiration: time.Hour,
	})
}
