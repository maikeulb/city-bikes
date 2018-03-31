package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/cache"
	"github.com/gorilla/mux"
	"github.com/maikeulb/city-bikes/redis"
)

type App struct {
	Router *mux.Router
}

func (a *App) InitializeServer() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("/api/networks")
	fmt.Println("/api/networks/{id}")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/networks", a.getNetworks).Methods("GET")
	a.Router.HandleFunc("/api/networks/{id}", a.getNetwork).Methods("GET")
}

func (a *App) getNetworks(w http.ResponseWriter, r *http.Request) {
	cacheKey := "all_networks"
	log.Println("cache key - ", cacheKey)

	var networks NetworksResponse
	startQuery := time.Now()
	if err := redis.Codec.Get(cacheKey, &networks); err != nil {
		response, err := http.Get("http://api.citybik.es/v2/networks")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		serializedNetworks, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(serializedNetworks, &networks)

		updateCacheNetworksResponse(cacheKey, networks)
		endQuery := time.Now()
		log.Println("retrieved networks from remote api in ", endQuery.Sub(startQuery), " seconds")
	} else {
		endQuery := time.Now()
		log.Println("retrieved networks from cache in ", endQuery.Sub(startQuery), " seconds")
	}

	respondWithJSON(w, http.StatusOK, networks)
}

func (a *App) getNetwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	cacheKey := id
	log.Println("cache key - ", cacheKey)

	var network NetworkResponse
	startQuery := time.Now()
	if err := redis.Codec.Get(cacheKey, &network); err != nil {

		response, err := http.Get("http://api.citybik.es/v2/networks/" + id)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		serializedNetwork, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(serializedNetwork, &network)

		updateCacheNetworkResponse(cacheKey, network)
		endQuery := time.Now()
		log.Println("retrieved network from remote api in ", endQuery.Sub(startQuery), " seconds")
	} else {
		endQuery := time.Now()
		log.Println("retrieved network from cache in ", endQuery.Sub(startQuery), " seconds")
	}

	respondWithJSON(w, http.StatusOK, network)
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

func updateCacheNetworksResponse(key string, networks NetworksResponse) {
	redis.Codec.Set(&cache.Item{
		Key:        key,
		Object:     networks,
		Expiration: time.Hour,
	})
}

func updateCacheNetworkResponse(key string, network NetworkResponse) {
	redis.Codec.Set(&cache.Item{
		Key:        key,
		Object:     network,
		Expiration: time.Hour,
	})
}
