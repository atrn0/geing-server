package http

import (
	"github.com/aratasato/geing-server/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	db                  *db.Conn
	adminUser           *string
	adminPass           *string
	netlifyBuildHookURL *string
	serverBaseUrl       *string
	corsAllowOrigin     *string
}

func NewServer(
	db *db.Conn,
	adminUser,
	adminPass,
	netlifyBuildHookURL,
	serverBaseUrl,
	corsAllowOrigin *string,
) *Server {
	return &Server{
		db,
		adminUser,
		adminPass,
		netlifyBuildHookURL,
		serverBaseUrl,
		corsAllowOrigin,
	}
}

func (s *Server) Start() error {
	router := s.Routes()
	return http.ListenAndServe(":9090", router)
}

func (s *Server) setHeader(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", *s.corsAllowOrigin)
		h(w, r, ps)
	}
}

func basicAuth(h httprouter.Handle, requiredUser, requiredPassword *string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == *requiredUser && password == *requiredPassword {
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
	router.GET("/questions", s.setHeader(s.getQuestions))
	router.GET("/questions/:uid", s.setHeader(s.getQA))
	router.POST("/questions", s.setHeader(s.addQuestion))
	router.GET("/admin", basicAuth(s.admin, s.adminUser, s.adminPass))
	router.GET("/admin/answer/:uid", basicAuth(s.getAnswerForm, s.adminUser, s.adminPass))
	router.POST("/admin/answer/:uid", basicAuth(s.addAnswer, s.adminUser, s.adminPass))
	return router
}
