package main

import (
	"log"
	"net/http"
	
	"goserver/internal/router"
	"goserver/internal/database"
)

func main() {
	// Инициализация базы данных
	database.InitDB("custom_data.db")

	// Создание маршрутизатора
	r := router.NewRouter()

	// Запуск сервера
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}






