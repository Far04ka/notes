package main

import (
	"net/http"
	"notes/constants"
	"notes/pages"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/auth", pages.Auth_form)
	mux.HandleFunc("/logout", pages.Quit)
	mux.HandleFunc("/new_note", pages.NewNote)
	mux.HandleFunc("/note", pages.Note)
	mux.HandleFunc("/note_act", pages.DelRedNote)
	mux.HandleFunc("/", pages.Route)

	http.ListenAndServe(constants.ServerPath, mux)
}
