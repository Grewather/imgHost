package main

import (
	"context"
	"imgHost/db"
	"imgHost/route"
	"net/http"
)

func main() {
	//r := chi.Router()
	db.ConnectToDb()
	defer db.Client.Disconnect(context.TODO())
	r := route.Router()

	http.ListenAndServe(":3000", r)

	//var r *chi.Mux = chi.Router()
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello, world!"))
	//
	//})
}
