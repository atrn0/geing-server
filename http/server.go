package http

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/questions", s.getQuestions)
	router.GET("/questions/:uid", s.getQA)
	router.POST("/questions", s.addQuestion)
	log.Fatalln(http.ListenAndServe(":9090", router))

	return router
}
