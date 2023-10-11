package Database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // SQL Package
)

func Connect() (*sql.DB, error) {
	connStr := "user=patoanctkz password=R7cvUjxB3uTV dbname=neondb host=ep-round-credit-135823.ap-southeast-1.aws.neon.tech sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
