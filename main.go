package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	. "questionBoxWithGo/models"
	"questionBoxWithGo/repository"
	. "questionBoxWithGo/utils"
	"strconv"
)

// 質問を20件取得
func getQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		// TODO: エラー処理を抽象化
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("page is invalid"))
		HandleError(err)
		return
	}

	res, err := repository.GetQuestions(page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write(ErrRes("questions not found"))
		HandleError(err)
		return
	}

	b, err := json.Marshal(res)
	HandleError(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

	_, err = w.Write(b)
	HandleError(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}
}

// 個別の質問と回答
func getQA(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	uidStr := ps.ByName("uid")

	// 数字かどうか
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("question id is invalid"))
		HandleError(err)
		return
	}

	res, err := repository.GetQA(uid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

	// TODO: 回答がまだなかったら404

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}
}

// 質問を投稿
func addQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	questionBody := r.FormValue("body")
	if questionBody == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(ErrRes("body is required"))
		HandleError(err)
		return
	}

	if repository.CreateQuestion(questionBody) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes("database error"))
		HandleError(err)
		return
	}

	// TODO: 実際の値を返す
	res := AddQuestionsResponse{0, questionBody, "Wed Oct 23 2019 12:56:05 GMT+0900"}
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(ErrRes(err.Error()))
		HandleError(err)
		return
	}

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
