package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type AddQuestionsResponse struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

type QA struct {
	Id        int    `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt string `json:"created_at"`
}

// 質問を20件取得
func getQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	page := r.URL.Query().Get("page")
	_, err = strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(errRes("page is invalid"))
		handleError(err)
		return
	}

	res := []QA{
		{
			0,
			"this is question",
			"this is answer",
			"Wed Oct 23 2019 12:56:05 GMT+0900",
		},
		{
			1,
			"this is question",
			"this is answer",
			"Wed Oct 23 2019 12:56:05 GMT+0900",
		},
	}
	b, err := json.Marshal(res)
	handleError(err)

	_, err = w.Write(b)
	handleError(err)
}

// 個別の質問と回答
func getQA(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var err error
	uid := ps.ByName("uid")

	i, err := strconv.Atoi(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(errRes("question id is invalid"))
		handleError(err)
		return
	}

	res := QA{i, "this is question", "this is answer", "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	handleError(err)

	i, err = w.Write(b)
	handleError(err)
}

// 質問を投稿
func addQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	questionBody := r.FormValue("body")
	if questionBody == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(errRes("body is required"))
		handleError(err)
		return
	}

	res := AddQuestionsResponse{0, questionBody, "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	handleError(err)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	handleError(err)
}

func main() {
	router := httprouter.New()
	router.GET("/questions", getQuestions)
	router.GET("/questions/:uid", getQA)
	router.POST("/questions", addQuestion)
	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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
