package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"questionBoxWithGo/db"
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

func (s *Server) Start(port string) error {
	router := s.Routes()
	fmt.Println("Listening on " + port)
	return http.ListenAndServe(":"+port, router)
}
