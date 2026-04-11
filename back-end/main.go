package main

import (
	"fullstack/db"
	"fullstack/route"
	"log"
	"net/http"
)

func main() {
	db.Init()
	route.SetupRouter()
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
