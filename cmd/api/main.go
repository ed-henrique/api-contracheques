package main

import (
	"api-contracheques/internal"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type server struct {
  db *sql.DB
  server *http.ServeMux
}

const (
  PORT = 8080
)

func (s *server) EmployeeById(w http.ResponseWriter, r *http.Request) {
  idPath := r.PathValue("id")

  id, err := strconv.Atoi(idPath)

  if err != nil {
    http.Error(w, "Por favor, insira um id válido (inteiro positivo não-nulo)", http.StatusBadRequest)
    return
  }

  fmt.Fprintf(w, "Funcionário com o id %d alcançado", id) 
}

func main() {
  db, err := internal.NewDatabase()

  if err != nil {
    log.Fatalf("Não foi possível iniciar o banco de dados: %s", err)
  }

  s := &server{
    db: db,
    server: http.NewServeMux(),
  }

  // Rotas
  s.server.HandleFunc("GET /funcionarios/{id}", s.EmployeeById)
  s.server.HandleFunc("POST /funcionarios/{id}", func(w http.ResponseWriter, r *http.Request) {})
  s.server.HandleFunc("GET /funcionarios/{id}/contracheque", func(w http.ResponseWriter, r *http.Request) {})

  if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), s.server); err != nil {
    log.Fatalf("Não foi possível iniciar um servidor na porta no endereço :%d\n", PORT)
  }
}
