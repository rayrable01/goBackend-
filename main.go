package main

import (
	"net/http"
	"secondProject/notes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	store := notes.NewNoteStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/notes", func(r chi.Router) {
		r.Get("/", notes.GetAllNotes(store))
		r.Post("/", notes.CreateNote(store))
		r.Get("/{id}", notes.GetNoteByID(store))
		r.Delete("/{id}", notes.DeleteNoteByID(store))
	})

	http.ListenAndServe(":8080", r)
}
