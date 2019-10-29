package http

import (
	"github.com/aratasato/geing-server/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	db *db.Conn
}

func NewServer(db *db.Conn) *Server {
	return &Server{db}
}

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/questions", s.getQuestions)
	router.GET("/questions/:uid", s.getQA)
	router.POST("/questions", s.addQuestion)

	return router
}

func (s *Server) Start() error {
	router := s.Routes()
	return http.ListenAndServe(":9090", router)
}
