package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sound/random", getRandomSound)
	router.HandleFunc("/sound/random/{genre}", getRandomSoundByGenre)

	fmt.Println("API runing at port 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}

type Sounds struct {
	Sounds []Sound `json:"sounds"`
}

type Sound struct {
	Music  string `json:"music"`
	Genre  string `json:"genre"`
	Movie  string `json:"movie"`
	Singer string `json:"verse_singer"`
	Media  *Media `json:"media"`
}

type Media struct {
	Sound   string   `json:"sound"`
	Picture *Picture `json:"picture"`
}

type Picture struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

func getRandomSound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonFile, err := os.Open("database/sounds.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var sounds Sounds

	w.WriteHeader(http.StatusCreated)

	json.Unmarshal(byteValue, &sounds)

	randomSound := rand.Int() % len(sounds.Sounds)

	fmt.Println(randomSound)

	json.NewEncoder(w).Encode(sounds.Sounds[randomSound])
}

func getRandomSoundByGenre(w http.ResponseWriter, r *http.Request) {

}
