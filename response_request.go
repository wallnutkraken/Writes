package main

type Response struct {
	OK bool
	Value interface{}
}

func RespNoSuchNote(triedId string) Response {
	return Response{OK: false, Value: "Could not find note with ID " + triedId}
}