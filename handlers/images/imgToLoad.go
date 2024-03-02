package images

import (
	"encoding/json"
	"fmt"
	"imgHost/db"
	"imgHost/utils"
	"net/http"
)

func ImgToLoad(w http.ResponseWriter, r *http.Request) {
	accessToken, err := r.Cookie("token")
	userInfo, err := utils.GetUserInfo(accessToken.Value)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	imagesFirstPage, err := db.GetImagesToLoad(userInfo.DiscordId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	urls := make([]string, len(imagesFirstPage))
	for i, image := range imagesFirstPage {
		urls[i] = image.Url
	}

	urlsJSON, err := json.Marshal(urls)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(urlsJSON)
}
