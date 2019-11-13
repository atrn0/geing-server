package http

import (
	"encoding/json"
	"fmt"
	"github.com/aratasato/geing-server/db"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Msg string `json:"msg"`
}

type GetQuestionsResponse struct {
	Questions []db.Question `json:"questions"`
}

type AddQuestionsResponse struct {
	QuestionBody string `json:"question_body"`
}

type AddQuestionsRequest struct {
	Body string `json:"body"`
}

type GetAnswerViewRequest struct {
	Id       int
	Question string
	Answer   *string
}

type AddAnswerRequest struct {
	Body string `json:"body"`
}

// TODO: headerを共通化

// 質問を20件取得
func (s *Server) getQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	var res []byte

	fmt.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pageStr := r.URL.Query().Get("page")
	// pageの値がなかったら最初のページを返す
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Println(err)
		msg := "page is invalid"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	questions, err := s.db.GetQuestions(page)
	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	if questions == nil {
		msg := "questions not found"
		w.WriteHeader(http.StatusNotFound)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	res, _ = json.Marshal(GetQuestionsResponse{questions})
	_, _ = w.Write(res)
	fmt.Println("res: ", string(res))
}

// 個別の質問と回答
func (s *Server) getQA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var res []byte
	var err error

	fmt.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	uidStr := ps.ByName("uid")

	// 数字かどうか
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		fmt.Println(err)
		msg := "question id should be integer"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	qa, err := s.db.GetQA(uid)

	if err == db.ErrContentNotFound {
		w.WriteHeader(http.StatusNotFound)
		res, _ = json.Marshal(ErrorResponse{"question not found"})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ = json.Marshal(qa)
	_, _ = w.Write(res)
	fmt.Println("res: ", string(res))
}

// 質問を投稿
func (s *Server) addQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var res []byte
	var err error

	fmt.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Header.Get("Content-type") != "application/x-www-form-urlencoded" {
		msg := "invalid Content-type"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	var reqBody AddQuestionsRequest
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		msg := "invalid request"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	questionBody := reqBody.Body
	if questionBody == "" {
		msg := "question is required"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	err = s.db.SaveQuestion(questionBody)
	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	newQuestion := AddQuestionsResponse{questionBody}
	res, _ = json.Marshal(newQuestion)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
	fmt.Println("res: ", string(res))
}

// 質問に回答する用のviewを返す
func (s *Server) getAnswerForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var res []byte

	t, err := template.ParseFiles("view/answer.html")
	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	uidStr := p.ByName("uid")

	// 数字かどうか
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		fmt.Println(err)
		msg := "question id should be integer"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	qa, err := s.db.GetQA(uid)
	if err == db.ErrContentNotFound {
		w.WriteHeader(http.StatusNotFound)
		res, _ = json.Marshal(ErrorResponse{"question not found"})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	_ = t.Execute(w, GetAnswerViewRequest{qa.Id, qa.Question, qa.Answer})
}

// 質問に回答
func (s *Server) addAnswer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var res []byte
	var err error

	fmt.Println(r.Method, r.URL)

	if r.Header.Get("Content-type") != "application/x-www-form-urlencoded" {
		msg := "invalid Content-type"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	err = r.ParseForm()
	if err != nil {
		msg := "invalid request"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	answerBody := r.Form["body"][0]
	if answerBody == "" {
		msg := "answer is required"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	uidStr := p.ByName("uid")

	// 数字かどうか
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		fmt.Println(err)
		msg := "question id should be integer"
		w.WriteHeader(http.StatusBadRequest)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	err = s.db.SaveAnswer(answerBody, uid)
	if err != nil {
		fmt.Println(err)
		msg := "internal server error"
		w.WriteHeader(http.StatusInternalServerError)
		res, _ = json.Marshal(ErrorResponse{msg})
		_, _ = w.Write(res)
		fmt.Println("res: ", string(res))
		return
	}

	http.Redirect(w, r, r.RequestURI, 301)
}
