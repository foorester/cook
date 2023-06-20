package http

import (
	"context"
	"net/http"
)

type (
	ContextKey string
)

const (
	BookCtxKey = "book"
)

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID := "a4e52de2-352b-4a61-964b-7efb7c137538" // WIP: Extract it from path

		ctx := context.WithValue(r.Context(), BookCtxKey, bookID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
