package images

import (
	"imgHost/utils"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	accessToken, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userinfo, err := utils.GetUserInfo(accessToken.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if _, err := os.Stat("images/" + userinfo.ID); os.IsNotExist(err) {
		os.MkdirAll("images/"+userinfo.ID, 0755)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	outFile, err := os.Create("images/" + userinfo.ID + "/" + fileHeader.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("File uploaded successfully"))
}
