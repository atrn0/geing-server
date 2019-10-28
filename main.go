package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"questionBoxWithGo/db"
	. "questionBoxWithGo/models"
	. "questionBoxWithGo/utils"
	"strconv"
)

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
