package admin

import (
	"imgHost/db"
	"net/http"
)

func AddInv(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invitedId := r.FormValue("invitedId")
	result, err := db.AddInvitedId(invitedId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
