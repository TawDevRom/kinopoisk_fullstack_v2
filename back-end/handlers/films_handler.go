package handlers

import (
	"encoding/json"
	"fullstack/db"
	"fullstack/models"
	"log"
	"net/http"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	query := `SELECT 
		f.film_id,
		f.title,
		f.is_serial,
		COALISCE(f.description, ''),
		t.trailer_id,
		t.path,
		fc.film_card_id,
		fc.film_id,
		fc.path,
		fc.is_horizontal,
		l.logo_id,
		l.path
	FROM films f
	LEFT JOIN trailers t 	 ON t.trailer_id = f.trailer_id
	LEFT JOIN film_cards fc  ON fc.film_id = f.film_id
	LEFT JOIN logos_films lf ON lf.film_id = f.film_id
	LEFT JOIN logos l 		 ON l.logo_id = lf.logo_id
	`


	rows, err := db.DB.Query(query)
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
