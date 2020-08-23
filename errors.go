package main

import "net/http"

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Err(message string, status int) *Error {
	return &Error{message, status}
}

var DatabaseError = Err("Your request failed due to a database error.", http.StatusInternalServerError)
