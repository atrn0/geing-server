package utils

import (
	"encoding/json"
	"log"
)

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ErrRes(msg string) []byte {
	b, err := json.Marshal(ErrorResponse{msg})
	HandleError(err)
	return b
}
