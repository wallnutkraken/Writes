package main

import (
	"encoding/json"
	"github.com/satori/go.uuid"
)

type note struct {
	Name string
	Content string
	Id string
}

func NewNote(name string) Note {
	n := new(note)
	n.Name = name
	n.Id = uuid.NewV4().String()

	return n
}

func (n *note) CopyOf() Note {
	newNote := NewNote("Copy of " + n.Name)
	newNote.SetContent(n.Content)
	return newNote
}

func (n *note) ToJSON() []byte {
	jsonBytes, _ := json.Marshal(n)
	return jsonBytes
}

func (n *note) GetContent() string {
	return n.Content
}

func (n *note) SetContent(content string) {
	n.Content = content
}

func (n *note) AppendContent(content string) {
	n.Content += content
}

func (n *note) GetId() string {
	return n.Id
}

func (n *note) GetName() string {
	return n.Name
}

func (n *note) Rename(name string) {
	n.Name = name
}

func JSONToNote(jsonBytes []byte) (Note, error) {
	var n *note
	n = new(note)
	err := json.Unmarshal(jsonBytes, n)
	return n, err
}

type Note interface {
	JSONObject
	GetContent() string
	SetContent(string)
	AppendContent(string)
	GetId() string
	CopyOf() Note
	GetName() string
	Rename(string)
}