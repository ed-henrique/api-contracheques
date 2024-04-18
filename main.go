package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetFuncionarioById(w http.ResponseWriter, r *http.Request) {
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
  server.HandleFunc("GET /funcionarios/{id}", GetFuncionarioById)
  server.HandleFunc("POST /funcionarios/{id}", func(w http.ResponseWriter, r *http.Request) {})
  server.HandleFunc("GET /funcionarios/{id}/contracheque", func(w http.ResponseWriter, r *http.Request) {})
}
