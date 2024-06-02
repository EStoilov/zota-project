package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	loggingHandler := LoggingMiddleware(handler)
	start := time.Now()
	loggingHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if duration := time.Since(start); duration <= 0 {
		t.Errorf("expected positive duration, got %v", duration)
	}
}