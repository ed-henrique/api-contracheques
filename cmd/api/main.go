package main

import (
	"api-contracheques/internal"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type server struct {
  db *sql.DB
  server *http.ServeMux
}

const (
  PORT = 8080
  DB_PATH = "local.db"
)

func main() {
  db, err := internal.NewDatabase(DB_PATH)

  if err != nil {
    log.Fatalf("It was not possible to start the database: %s", err)
  }

  s := &server{
    db: db,
    server: http.NewServeMux(),
  }

  // Rotas
  s.server.HandleFunc("GET /employees/{id}", s.employeeById)
  s.server.HandleFunc("POST /employees/{id}", func(w http.ResponseWriter, r *http.Request) {})
  s.server.HandleFunc("GET /employees/{id}/paycheck", func(w http.ResponseWriter, r *http.Request) {})

  if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), s.server); err != nil {
    log.Fatalf("It was not possible to start the server in the address :%d\n", PORT)
  }
}
