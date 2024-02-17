package images

import (
	"github.com/go-chi/chi/v5"
	"imgHost/db"
	"net/http"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ownerId, imgName, err := db.GetImg(id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	http.ServeFile(w, r, "uploads/"+ownerId+"/"+imgName)
}
