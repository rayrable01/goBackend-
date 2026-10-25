package notes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAllNotes(store *NoteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes := store.GetAll()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func CreateNote(store *NoteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var noteDTO Note

		if err := json.NewDecoder(r.Body).Decode(&noteDTO); err != nil {
			errorDTO := ErrorResponse{
				Error:   err.Error(),
				Message: "Can't read properties of JSON",
			}

			http.Error(w, errorDTO.Message, http.StatusBadRequest)
			return
		}

		note := store.Create(noteDTO.Title, noteDTO.Content)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(note); err != nil {
			return
		}
	}
}

func GetNoteByID(store *NoteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || noteID <= 0 {
			http.Error(w, "Invalid note ID", http.StatusBadRequest)
			return
		}

		note, ok := store.GetByID(noteID)
		if !ok {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(note); err != nil {
			return
		}
	}
}

func DeleteNoteByID(store *NoteStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil || noteID <= 0 {
			http.Error(w, "Invalid note ID", http.StatusBadRequest)
			return
		}

		ok := store.Delete(noteID)

		if !ok {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ok); err != nil {
			return
		}
	}
}
