package main

import (
	"fmt"
	"github.com/aratasato/geing-server/db"
	"github.com/aratasato/geing-server/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	fmt.Println("Starting geing server...")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// init conn
	conn, err := db.NewDB()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("init db")

	// init and start server
	server := http.NewServer(
		conn,
		os.Getenv("ADMIN_USERNAME"),
		os.Getenv("ADMIN_PASSWD"),
	)
	fmt.Println("init server")
	err = server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
