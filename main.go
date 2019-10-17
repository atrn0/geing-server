package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var itemList = []Item{
	Item{Title: "Item A", Description: "The first item"},
	Item{Title: "Item B", Description: "The second item"},
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page.")
}

func returnAllItems(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, itemList)
}

func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "the payload is invalid")
		return
	}

	if key >= len(itemList) {
		respondWithError(w, http.StatusNotFound, "Item does not exists")
		return
	}

	respondWithJson(w, http.StatusOK, itemList[key])
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/items", returnAllItems)
	myRouter.HandleFunc("/{id}", returnSingleItem)
	log.Fatal(http.ListenAndServe(":8888", myRouter))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	handleRequests()
}
