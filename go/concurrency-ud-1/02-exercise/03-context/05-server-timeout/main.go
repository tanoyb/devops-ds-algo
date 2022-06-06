package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func slowQuery(ctx context.Context) error {
	_, err := db.ExecContext(ctx, "SELECT pg_sleep(5)")
	return err
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	err := slowQuery(req.Context())
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Fprintln(w, "OK")
	fmt.Printf("slowHandler took: %v\n", time.Since(start))
}

func main() {
	var err error

	connstr := "host=localhost port=5432 user=review_admin password=review_admin  dbname=review_db_uqus_stg sslmode=disable"

	db, err = sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:         "localhost:8000",
		WriteTimeout: 2 * time.Second,
		Handler: http.TimeoutHandler(
			http.HandlerFunc(slowHandler),
			1*time.Second,
			"timeout!",
		),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

// --> Installing postgres - macos
// brew install postgresql

// --> start
// pg_ctl -D /usr/local/var/postgres start

// --> create db and user
// psql postgres
// CREATE DATABASE wonderland;
// CREATE USER alice WITH ENCRYPTED PASSWORD 'pa$$word';
// GRANT ALL PRIVILEGES ON DATABASE wonderland TO alice;

// --> stop
// pg_ctl -D /usr/local/var/postgres stop

// --> postgresql download link
// https://www.postgresql.org/download/

// start postgresql - Windows
// pg_ctl -D "C:\Program Files\PostgreSQL\13\data" start

// stop postgresql - Windows
// pg_ctl -D "C:\Program Files\PostgreSQL\13\data" stop

// --> Linux
// sudo apt-get update
// sudo apt-get install postgresql-13

// sudo -u postgres psql -c "ALTER USER alice PASSWORD 'pa$$word';"
// sudo -u postgres psql -c "CREATE DATABASE wonderland;"

// sudo service postgresql start

// sudo service postgresql stop
