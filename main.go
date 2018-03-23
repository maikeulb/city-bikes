package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Networks []Network `json:"networks"`
}

type Network struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	City      int     `json:"city"`
	Country   int     `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
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

	for i := 0; i < len(responseObject.Networks); i++ {
		fmt.Println(responseObject.Networks[i].Name)
	}
}
