package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	payload := map[string]string{"key": "value"}
	resp, err := SendPostRequest(server.URL, payload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", resp.Status)
	}
}