package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=Tawern228 dbname=online_kino sslmode=disable"
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
	query := `
		CREATE TABLE IF NOT EXISTS films (
			film_id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			is_serial BOOLEAN NOT NULL DEFAULT false,
			description TEXT,
			trailer_id INT
		);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Ошибка миграции", err)
	}
	log.Println("Таблица films готова")
}
