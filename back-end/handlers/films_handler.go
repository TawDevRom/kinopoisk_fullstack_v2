package handlers

import (
	"encoding/json"
	"fullstack/db"
	"fullstack/models"
	"log"
	"net/http"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT film_id, title, is_serial, description FROM films;")
	if err != nil {
		log.Fatal("Ошибка запроса")
		return
	}
	defer rows.Close()

	var films []models.Films

	for rows.Next() {
		var f models.Films
		err := rows.Scan(&f.ID, &f.Title, &f.IsSerial, &f.Description)
		if err != nil {
			log.Println("Ошибка чтения строки", err)
			continue
		}
		films = append(films, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}
