package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *server) employeeById(w http.ResponseWriter, r *http.Request) {
  idPath := r.PathValue("id")

  id, err := strconv.Atoi(idPath)

  if err != nil || id <= 0 {
    http.Error(w, `{"error":"invalid employee id (should be a non-null integer)"}`, http.StatusBadRequest)
    return
  }

  fmt.Fprintf(w, "Funcionário com o id %d alcançado", id) 
}
