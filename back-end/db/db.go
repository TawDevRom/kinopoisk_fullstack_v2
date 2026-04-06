package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=online_kino sslmode=disable" //Tawern228
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключение к БД", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("БД не отвечает", err)
	}

	log.Println("Подключение к БД успешно")
	migrate()
}

func migrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS trailers (
			trailer_id SERIAL PRIMARY KEY,
			path TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS logos (
			logo_id SERIAL PRIMARY KEY,
			path TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS films (
			film_id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			is_serial BOOLEAN NOT NULL DEFAULT false,
			description TEXT,
			trailer_id INT REFERENCES trailers(trailer_id)
		);`,
		`CREATE TABLE IF NOT EXISTS film_cards (
			film_card_id SERIAL PRIMARY KEY,
			film_id INT REFERENCES films(film_id),
			path TEXT NOT NULL,
			is_horizontal BOOLEAN NOT NULL DEFAULT true
		);`,
		`CREATE TABLE IF NOT EXISTS logos_films (
			logo_film_id SERIAL PRIMARY KEY,
			logo_id INT REFERENCES logos(logo_id),
			film_id INT REFERENCES films(film_id)
		);`,
	}
	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Ошибка миграции", err)
		}
	}

	log.Println("Миграция выполнена успешно")
}

// INSERT INTO films (title, is_serial, description) VALUES ('Атака титанов', true, 'Люди сражаются с титанами, которые мечтают их съесть. Финал самого эпичного аниме современности'),
// ('Дандадан', true, 'Внучка медиума и юный уфолог внезапно встречают призраков и пришельцев. Хитовое аниме - безумное и смешное'),
// ('Железный человек', false, 'Он не просто металл - он живой! Идём в кино на высокотехнологичную легенду, соединяющую поколения.');
