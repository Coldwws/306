package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Coldwws/306/internal/handler"
	"github.com/Coldwws/306/internal/repository"
	"github.com/Coldwws/306/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal("Ошибка подключения к бд", err)
	}
	storage := repository.NewRoomPostgres(db)
	roomService := service.NewRoomImplementation(storage)

	roomHandler := handler.NewRoomHandler(roomService)

	router := roomHandler.InitRoutes()

	server := http.Server{
		Addr:    ":5588",
		Handler: router,
	}

	fmt.Printf("Server started on %s", server.Addr)
	server.ListenAndServe()
}
