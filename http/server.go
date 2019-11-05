package http

import (
	"github.com/aratasato/geing-server/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	db *db.Conn
}

func setHeader(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h(w, r, ps)
	}
}

func NewServer(db *db.Conn) *Server {
	return &Server{db}
}

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/questions", setHeader(s.getQuestions))
	router.GET("/questions/:uid", setHeader(s.getQA))
	router.POST("/questions", setHeader(s.addQuestion))

	return router
}

func (s *Server) Start() error {
	router := s.Routes()
	return http.ListenAndServe(":9090", router)
}
