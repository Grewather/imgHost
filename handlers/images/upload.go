package images

import (
	"fmt"
	"imgHost/db"
	"imgHost/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	if _, err := os.Stat("uploads/" + userinfo.ID); os.IsNotExist(err) {
		os.MkdirAll("uploads/"+userinfo.ID, 0755)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	extension := filepath.Ext(header.Filename)
	randString := utils.GetRandomString()
	fmt.Println("3")

	for {
		if checkIfYouCanAdd(randString, userinfo.ID, extension) {
			break
		}
		randString = utils.GetRandomString()
	}
	defer file.Close()
	outFile, err := os.Create("uploads/" + userinfo.ID + "/" + randString + extension)
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
	fmt.Println("1")
	url := db.AddImgToDb(randString, userinfo.ID, extension)
	fmt.Println("2")
	http.Redirect(w, r, "/i/"+url, http.StatusFound)
}

func checkIfYouCanAdd(randString, discordid, extension string) bool {
	pathfile := "uploads/" + discordid + "/" + randString

	if _, err := os.Stat(pathfile); err == nil {
		return false
	} else if os.IsNotExist(err) {
		return true
	} else {
		// better handling of error
		return false
	}
}
