package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	. "questionBoxWithGo/models"
	. "questionBoxWithGo/utils"
	"strconv"
)

// 質問を20件取得
func getQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	page := r.URL.Query().Get("page")
	_, err = strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("page is invalid"))
		HandleError(err)
		return
	}

	res := []GetQAsResponse{
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
	HandleError(err)

	_, err = w.Write(b)
	HandleError(err)
}

// 個別の質問と回答
func getQA(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var err error
	uid := ps.ByName("uid")

	i, err := strconv.Atoi(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("question id is invalid"))
		HandleError(err)
		return
	}

	res := GetQAsResponse{i, "this is question", "this is answer", "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	HandleError(err)

	i, err = w.Write(b)
	HandleError(err)
}

// 質問を投稿
func addQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	questionBody := r.FormValue("body")
	if questionBody == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("body is required"))
		HandleError(err)
		return
	}

	//if repository.CreateQuestion(questionBody) != nil {
	//	w.WriteHeader(http.)
	//}

	res := AddQuestionsResponse{0, questionBody, "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	HandleError(err)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	HandleError(err)
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
