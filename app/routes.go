package app

import (
	"github.com/vladvelici/quizz/infra"
	"github.com/vladvelici/quizz/user"

	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() {
	r := mux.NewRouter().StrictSlash(false)

	r.Handle("/", infra.Mid(HomePage, user.Auth))

	account := r.Path("/account").Subrouter()
	account.Methods("GET").Handler(infra.Mid(user.AccountGet, user.MustAuth))
	account.Methods("POST").Handler(infra.Mid(user.AccountPost, user.MustAuth))

	http.Handle("/", r)
}
