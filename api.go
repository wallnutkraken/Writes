package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"strconv"
)

type api struct {
	router *mux.Router
	notes []Note
}

func (a *api) GetNote(id string) Note {
	for _, n := range a.notes {
		if n.GetId() == id {
			return n
		}
	}
	return nil
}

func (a *api) startApi(port int) error {
	a.notes = make([]Note, 0)

	a.notes = append(a.notes, NewNote("testNote"))

	a.router = mux.NewRouter().StrictSlash(true)
	a.router.HandleFunc("/", Index)
	a.router.HandleFunc("/notes", a.NoteIndex)
	a.router.HandleFunc("/notes/{noteid}", a.TodoShow)

	return http.ListenAndServe(":" + strconv.Itoa(port), a.router)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "!")
	//todo: api info here
}

func (a *api) NoteIndex(w http.ResponseWriter, r *http.Request) {
	index := NoteIndex{}
	index.Count = len(a.notes)
	index.Notes = make([]NoteInfo, index.Count)
	for i, currNote := range a.notes {
		index.Notes[i] = NoteInfo{Name: currNote.GetName(), Id: currNote.GetId()}
	}
	w.Write(index.ToJSON())
}

func (a *api) TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["noteid"]
	n := a.GetNote(noteId)
	if n == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(RespNoSuchNote(noteId).ToJSON())
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(n.ToJSON())
	}
}