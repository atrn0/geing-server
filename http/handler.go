package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"questionBoxWithGo/db"
	"strconv"
)

type ErrorResponse struct {
	Msg string `json:"msg"`
}

type AddQuestionsResponse struct {
	QuestionBody string `json:"question_body"`
}

// TODO: Serverがdbを持つ

// 質問を20件取得
func (s *Server) getQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	var res []byte

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageStr := r.URL.Query().Get("page")
	// pageの値がなかったら最初のページを返す
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{"page is invalid"})
		_, _ = w.Write(res)
		return
	}

	questions, err := db.GetQuestions(page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{"internal server error"})
		_, _ = w.Write(res)
		return
	}

	if questions == nil {
		w.WriteHeader(http.StatusNotFound)
		res, _ = json.Marshal(ErrorResponse{"questions not found"})
		_, _ = w.Write(res)
		return
	}

	res, err = json.Marshal(questions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{"internal server error"})
		_, _ = w.Write(res)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{"internal server error"})
		_, _ = w.Write(res)
		return
	}
}

// 個別の質問と回答
func (s *Server) getQA(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	var res []byte
	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	uidStr := ps.ByName("uid")

	// 数字かどうか
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{"question id should be integer"})
		_, _ = w.Write(res)
		return
	}

	qa, err := db.GetQA(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{"internal server error"})
		_, _ = w.Write(res)
		return
	}

	// TODO: そのidのqaがなかったとき404

	w.WriteHeader(http.StatusOK)
	res, _ = json.Marshal(qa)
	_, _ = w.Write(res)
}

// 質問を投稿
func (s *Server) addQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var res []byte

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	questionBody := r.FormValue("body")
	if questionBody == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{"question is required"})
		_, _ = w.Write(res)
		return
	}

	if db.CreateQuestion(questionBody) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{"internal server error"})
		_, _ = w.Write(res)
		return
	}

	newQuestion := AddQuestionsResponse{questionBody}
	res, _ = json.Marshal(newQuestion)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}
