package handlers

import (
	"encoding/json"
	"net/http"
)

type Note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

// var notes []note
var notes = []Note{
	{ID: 1, Title: "Sample Note", Content: "This is a sample note"},
}

func HandleNotes(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(notes)
		return
	}

	if r.Method == http.MethodPost {
		var note Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		note.ID = len(notes) + 1
		notes = append(notes, note)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}