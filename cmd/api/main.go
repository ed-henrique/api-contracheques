package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
  PORT = 8080
)

func EmployeeById(w http.ResponseWriter, r *http.Request) {
  idPath := r.PathValue("id")

  id, err := strconv.Atoi(idPath)

  if err != nil {
    http.Error(w, "Por favor, insira um id válido (inteiro positivo não-nulo)", http.StatusBadRequest)
    return
  }

  fmt.Fprintf(w, "Funcionário com o id %d alcançado", id) 
}

func main() {
  server := http.NewServeMux()

  // Rotas
  server.HandleFunc("GET /funcionarios/{id}", EmployeeById)
  server.HandleFunc("POST /funcionarios/{id}", func(w http.ResponseWriter, r *http.Request) {})
  server.HandleFunc("GET /funcionarios/{id}/contracheque", func(w http.ResponseWriter, r *http.Request) {})

  if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), server); err != nil {
    log.Fatalf("Não foi possível iniciar um servidor na porta no endereço :%d\n", PORT)
  }
}
