package main

import (
	"api-contracheques/internal"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEmployeeById(t *testing.T) {
  t.Parallel()

  db, _ := internal.NewDatabase(":memory:")

  s := &server{
    db: db,
  }

  t.Run("fetch invalid employee", func(t *testing.T) {
    request := httptest.NewRequest(http.MethodGet, "/employees/0", nil)
    responseRecorder := httptest.NewRecorder()

    s.EmployeeById(responseRecorder, request)

    if responseRecorder.Code != http.StatusBadRequest {
      t.Errorf("got=%d expected=%d", responseRecorder.Code, http.StatusBadRequest)
    }

    if strings.TrimSpace(responseRecorder.Body.String()) != `{"error":"invalid employee id (should be a non-null integer)"}` {
      t.Errorf("got=%q expected=%q", responseRecorder.Body, `{"error":"invalid employee id (should be a non-null integer)"}`)
    }
  })

  t.Run("fetch non existent employee", func(t *testing.T) {
    request := httptest.NewRequest(http.MethodGet, "/employees/5000000", nil)
    responseRecorder := httptest.NewRecorder()

    s.EmployeeById(responseRecorder, request)

    if responseRecorder.Code != http.StatusNotFound {
      t.Errorf("got=%d expected=%d", responseRecorder.Code, http.StatusNotFound)
    }

    if strings.TrimSpace(responseRecorder.Body.String()) != `{"error":"employee not found"}` {
      t.Errorf("got=%q expected=%q", responseRecorder.Body, `{"error":"employee not found"}`)
    }
  })

  t.Run("fetch existing employee", func(t *testing.T) {
    request := httptest.NewRequest(http.MethodGet, "/employees/1", nil)
    responseRecorder := httptest.NewRecorder()

    s.EmployeeById(responseRecorder, request)

    if responseRecorder.Code != http.StatusOK {
      t.Errorf("got=%d expected=%d", responseRecorder.Code, http.StatusOK)
    }

    if strings.TrimSpace(responseRecorder.Body.String()) != `{"id":1,"name":"foo","surname":"bar","document":"00000000000","sector":"sales","gross_wage":'100000,"admission_date":"2024-01-01","has_healthcare":false,"has_dentalcare":true,"has_transportation_allowance":false}` {
      t.Errorf("got=%q expected=%q", responseRecorder.Body, `{"id":1,"name":"foo","surname":"bar","document":"00000000000","sector":"sales","gross_wage":'100000,"admission_date":"2024-01-01","has_healthcare":false,"has_dentalcare":true,"has_transportation_allowance":false}`)
    }
  })
}
