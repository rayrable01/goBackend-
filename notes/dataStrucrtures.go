package notes

import (
	"math/rand"
	"sync"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type NoteStore struct {
	mu     sync.RWMutex
	notes  map[int]Note
	nextID int
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func NewNoteStore() *NoteStore {
	newNoteStore := NoteStore{
		mu:     sync.RWMutex{},
		notes:  make(map[int]Note),
		nextID: 1,
	}

	return &newNoteStore
}

func (store *NoteStore) Create(
	title string,
	content string,
) Note {
	store.mu.Lock()
	defer store.mu.Unlock()

	newNote := Note{
		ID:        rand.Int(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	store.notes[newNote.ID] = newNote

	return newNote
}

func (store *NoteStore) GetAll() []Note {
	store.mu.RLock()
	defer store.mu.RUnlock()

	notes := make([]Note, 0, len(store.notes))

	for _, note := range store.notes {
		notes = append(notes, note)
	}

	return notes
}

func (store *NoteStore) GetByID(id int) (Note, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	val, ok := store.notes[id]

	if !ok {
		return Note{}, false
	}

	return val, true
}

func (store *NoteStore) Delete(id int) bool {
	store.mu.Lock()
	defer store.mu.Unlock()

	_, ok := store.notes[id]
	if !ok {
		return false
	}

	delete(store.notes, id)
	return true
}
