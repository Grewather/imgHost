package main

import (
	"context"
	"imgHost/db"
	"imgHost/route"
	"net/http"
)

func main() {
	db.ConnectToDb()
	defer db.Client.Disconnect(context.TODO())
	r := route.Router()

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}
