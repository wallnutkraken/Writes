package main

type NoteIndex struct {
	Count int
	Notes []NoteInfo
}

type NoteInfo struct {
	Name string
	Id string
}
