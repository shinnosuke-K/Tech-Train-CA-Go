package middleware

import (
	"net/http"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/handler/response"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/auth"
)

func POST(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func GET(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func PUT(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			response.Error(w, http.StatusMethodNotAllowed, nil, "bad request method")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		xToken := r.Header.Get("x-token")
		if xToken == "" {
			response.Error(w, http.StatusUnauthorized, nil, "x-token is empty")
			return
		}

		if err := auth.Validate(xToken); err != nil {
			response.Error(w, http.StatusUnauthorized, err, "x-token is invalid")
			return
		}
		next.ServeHTTP(w, r)
	}
}
