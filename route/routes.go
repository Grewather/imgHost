package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"html/template"
	"imgHost/handlers/auth"
	authMiddleware "imgHost/middleware"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	r.Use(corsMiddleware.Handler)
	r.Use(middleware.StripSlashes)
	r.Use(authMiddleware.AuthMiddleware)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./templates/index.html"))
			tmpl.Execute(w, nil)
		})
		r.Get("/upload", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./templates/upload.html"))
			tmpl.Execute(w, nil)
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API endpoint"))
		})

		r.Get("/auth/discord/login", auth.LoginAuth)
		r.Get("/auth/discord/callback", auth.LoginCallback)
	})

	return r
}
