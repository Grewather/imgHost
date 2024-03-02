package images

import (
	"fmt"
	"github.com/joho/godotenv"
	"imgHost/db"
	"imgHost/utils"
	"net/http"
	"os"
	"path/filepath"
)

type PostData struct {
	Secret string `json:"secret"`
	Source string `json:"source"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	secret := r.FormValue("secret")
	source := r.FormValue("source")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if source == "sharex" {
		userInfo, err := db.GetDataApiKey(secret)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		extension := filepath.Ext(header.Filename)
		randString := utils.GetRandomString()
		ext := utils.CheckFileExt(extension)
		if !ext {
			http.Error(w, "Invalid file extension", http.StatusBadRequest)
			return
		}

		err = utils.SaveFileToUploads(file, userInfo, extension, randString)
		if err != nil {
			if err.Error() == "invalid file extension" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		url := db.AddImgToDb(randString, userInfo.DiscordId, extension)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(os.Getenv("DOMAIN") + "/i/" + url))

	} else if source == "web" {
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

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		extension := filepath.Ext(header.Filename)
		fmt.Println(extension)
		randString := utils.GetRandomString()
		err = utils.SaveFileToUploads(file, userinfo, extension, randString)
		if err != nil {
			if err.Error() == "invalid file extension" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		url := db.AddImgToDb(randString, userinfo.DiscordId, extension)
		http.Redirect(w, r, "/i/"+url, http.StatusFound)
	} else {
		http.Error(w, "Invalid source", http.StatusBadRequest)
		return
	}

}
