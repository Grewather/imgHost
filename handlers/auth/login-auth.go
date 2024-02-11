package auth

import "net/http"

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	dscUrl := "https://discord.com/api/oauth2/authorize?client_id=1203287944115396648&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fapi%2Fauth%2Fdiscord%2Fcallback&scope=identify"
	http.Redirect(w, r, dscUrl, http.StatusMovedPermanently)
}
