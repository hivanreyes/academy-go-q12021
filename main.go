package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hivanreyes/academy-go-q12021/controller"
	"github.com/hivanreyes/academy-go-q12021/router"
	"github.com/hivanreyes/academy-go-q12021/service"
	"github.com/hivanreyes/academy-go-q12021/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Service
	s := service.New()
	u := usecase.New(s)
	c := controller.New(u)
	r := router.New(c)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
