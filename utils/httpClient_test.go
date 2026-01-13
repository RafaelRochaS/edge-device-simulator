package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakePostCall(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("response-body"))
		}))
		defer server.Close()

		resp, err := makePostCall(map[string]string{"key": "value"}, server.URL)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if resp != "response-body" {
			t.Errorf("expected response-body, got %s", resp)
		}
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		_, err := makePostCall(map[string]string{"key": "value"}, server.URL)
		if err == nil {
			t.Error("expected error for 500 status code")
		}
	})

	t.Run("invalid url", func(t *testing.T) {
		_, err := makePostCall(map[string]string{"key": "value"}, "http://invalid-url-that-does-not-exist")
		if err == nil {
			t.Error("expected error for invalid url")
		}
	})
}
