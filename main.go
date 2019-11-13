package main

import (
	"fmt"
	"github.com/aratasato/geing-server/db"
	"github.com/aratasato/geing-server/http"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	fmt.Println("Starting geing server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// init conn
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("init db")

	// init and start server
	server := http.NewServer(conn)
	fmt.Println("init server")
	err = server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
