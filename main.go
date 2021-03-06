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
	adminUser := os.Getenv("ADMIN_USERNAME")
	adminPass := os.Getenv("ADMIN_PASSWD")
	netlifyBuildHookURL := os.Getenv("NETLIFY_BUILD_HOOK_URL")
	serverBaseUrl := os.Getenv("SERVER_BASE_URL")
	corsAllowOrigin := os.Getenv("CORS_ALLOW_ORIGIN")
	iftttWebHookKey := os.Getenv("IFTTT_WEBHOOK_KEY")

	// init conn
	conn, err := db.NewDB()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("init db")

	// init and start server
	server := http.NewServer(
		conn,
		&adminUser,
		&adminPass,
		&netlifyBuildHookURL,
		&serverBaseUrl,
		&corsAllowOrigin,
		&iftttWebHookKey,
	)
	fmt.Println("init server")
	err = server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
