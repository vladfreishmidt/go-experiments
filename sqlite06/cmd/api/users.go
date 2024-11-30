package main

import "net/http"

func (app *application) showUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showing all users..."))
}
