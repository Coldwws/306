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
	if err := godotenv.Load(); err != nil {
		log.Print("не найден, использую переменную окружения из docker.")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal("Ошибка подключения к бд", err)
	}
	storage := repository.NewRoomPostgres(db)
	userRepo := repository.NewUserPostgres(db)

	roomService := service.NewRoomImplementation(storage)
	authService := service.NewAuthService(userRepo)

	roomHandler := handler.NewRoomHandler(roomService, authService)

	router := roomHandler.InitRoutes()

	server := http.Server{
		Addr:    ":5050",
		Handler: router,
	}

	fmt.Printf("Server started on %s", server.Addr)
	server.ListenAndServe()
}
