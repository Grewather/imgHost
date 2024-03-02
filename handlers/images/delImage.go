package images

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"imgHost/db"
	"imgHost/utils"
	"net/http"
	"os"
)

type DeleteResponse struct {
	Result bool `json:"result"`
}

func DeleteImg(w http.ResponseWriter, r *http.Request) {
	imgId := chi.URLParam(r, "id")
	accessToken, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Access token not found", http.StatusUnauthorized)
		return
	}
	userinfo, err := utils.GetUserInfo(accessToken.Value)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	imgToDelete, ext, result := db.DelFromDb(imgId)
	if !result {
		http.Error(w, "Image not found or user not authorized", http.StatusNotFound)
		return
	}
	response := DeleteResponse{
		Result: result,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = os.Remove("uploads/" + userinfo.DiscordId + "/" + imgToDelete + ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
