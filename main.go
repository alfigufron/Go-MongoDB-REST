package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/hexa/go-boilerplate-restapi/routes"
	"github.com/hexa/go-boilerplate-restapi/utils"
	"github.com/joho/godotenv"
)

func main() {
	clearConsole()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error .env file")
	}

	utils.InitLogger()
	utils.InitMongoDB()

	router := routes.InitRouter()

	port := os.Getenv("APP_PORT")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed Start HTTP Server!")
		log.Fatal(err)
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
