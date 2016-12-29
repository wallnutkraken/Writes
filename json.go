package main

import "encoding/json"

type JSONObject interface {
	ToJSON() []byte
}


func (n NoteIndex) ToJSON() []byte {
	jsonData, _ := json.Marshal(&n)
	return jsonData
}


func (r Response) ToJSON() []byte {
	jsonBytes, _ := json.Marshal(&r)
	return jsonBytes
}