package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/Go-MongoDB-REST/database"
	"github.com/Go-MongoDB-REST/routes"
	"github.com/Go-MongoDB-REST/utils"
	"github.com/joho/godotenv"
)

func main() {
	clearConsole()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error .env file")
	}

	utils.InitLogger()
	database.InitMongoDB()

	router := routes.InitRouter()

	port := os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		log.Fatal("Failed Start HTTP Server!")
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
