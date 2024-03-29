package route

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"imgHost/db"
	"imgHost/handlers/admin"
	"imgHost/handlers/auth"
	"imgHost/handlers/images"
	authMiddleware "imgHost/middleware"
	"imgHost/utils"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.StripSlashes)
	r.Use(authMiddleware.AuthMiddleware)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./templates/index.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		})
		r.Get("/upload", func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			userInfo, err := utils.GetUserInfo(cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			res, err := db.GetApiKey(userInfo.DiscordId)
			if err != nil {
				fmt.Println(err)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			userInfo.ApiKey = res
			data := struct {
				ApiKey string
			}{
				ApiKey: userInfo.ApiKey,
			}

			tmpl := template.Must(template.ParseFiles("./templates/upload.html"))
			err = tmpl.Execute(w, data)
			if err != nil {
				panic(err)
			}
		})
		r.Get("/gallery", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./templates/gallery.html"))
			err := tmpl.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		})

	})
	r.Route("/i", func(r chi.Router) {
		r.Get("/{id}", images.GetImage)
	})
	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/admin.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API endpoint"))
		})
		r.Post("/admin/addInvite", admin.AddInv)
		r.Post("/admin/removeInvite", admin.RemoveInv)
		r.Get("/images", images.ImgToLoad)
		r.Post("/upload", images.Upload)
		r.Delete("/delete/{id}", images.DeleteImg)
		r.Get("/auth/discord/logout", auth.Logout)
		r.Get("/auth/discord/login", auth.LoginAuth)
		r.Get("/auth/discord/callback", auth.LoginCallback)
	})

	return r
}
