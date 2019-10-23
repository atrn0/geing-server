package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type AddQuestionsResponse struct {
	Id        int32  `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func addQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	var res interface{}

	questionBody := r.FormValue("body")
	if questionBody == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write(errRes("body is required"))
		handleError(err)
		return
	}

	res = AddQuestionsResponse{0, questionBody, "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	handleError(err)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func errRes(msg string) []byte {
	b, err := json.Marshal(ErrorResponse{msg})
	handleError(err)
	return b
}

func main() {
	router := httprouter.New()
	router.POST("/questions", addQuestions)
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
