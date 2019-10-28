package main

import (
	"fmt"
	"log"
	"questionBoxWithGo/db"
	"questionBoxWithGo/http"
)

func main() {
	var port = "9090"

	fmt.Println("Starting geing server...")

	// init conn
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("init db")

	// init and start server
	server := http.NewServer(conn)
	fmt.Println("init server")
	err = server.Start(port)
	if err != nil {
		log.Fatalln(err)
	}
}
