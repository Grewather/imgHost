package admin

import (
	"imgHost/db"
	"net/http"
)

func RemoveInv(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invitedId := r.FormValue("invitedId")
	result := db.RemoveId(invitedId)
	if !result {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while removing id from whitelist"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Id removed from whitelist"))

}
