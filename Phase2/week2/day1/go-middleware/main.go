package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthMiddleware bertugas untuk mengecek apakah token kita valid atau tidak
func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		expectedToken := "secret-token"

		token := r.Header.Get("Authorization")

		if token != expectedToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, ps)
	}
}

func HandlerProtected(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Welcome Secret user"))
}

func HandlerPublic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Welcome Public user"))
}

func main() {
	router := httprouter.New()

	router.GET("/protected", AuthMiddleware(HandlerProtected))
	router.GET("/public", HandlerProtected)

	http.ListenAndServe(":8080", router)
}
