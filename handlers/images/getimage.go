package images

import (
	"github.com/go-chi/chi/v5"
	"imgHost/db"
	"net/http"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	ownerId, imgName, err := db.GetImg(id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	//w.Write([]byte("error"))
	http.ServeFile(w, r, "uploads/"+ownerId+"/"+imgName)
}
