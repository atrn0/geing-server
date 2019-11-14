package http

import (
	"github.com/aratasato/geing-server/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	db        *db.Conn
	adminUser string
	adminPass string
}

func NewServer(db *db.Conn, adminUser, adminPass string) *Server {
	return &Server{db, adminUser, adminPass}
}

func (s *Server) Start() error {
	router := s.Routes()
	return http.ListenAndServe(":9090", router)
}

func setHeader(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h(w, r, ps)
	}
}

func basicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func (s *Server) Routes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/questions", setHeader(s.getQuestions))
	router.GET("/questions/:uid", setHeader(s.getQA))
	router.POST("/questions", setHeader(s.addQuestion))
	router.GET("/admin/answer/:uid", basicAuth(s.getAnswerForm, s.adminUser, s.adminPass))
	router.POST("/admin/answer/:uid", basicAuth(s.addAnswer, s.adminUser, s.adminPass))
	return router
}
