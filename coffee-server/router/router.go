package router

import (
	"net/http"

	"github.com/Wexler763/CoffeeApiSecond/coffee-server/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*", "http://*", "https://*", "null"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/api/v1/cards", controllers.GetAllCards)
	router.Get("/api/v1/cards/thehornedcard/{id}", controllers.GetCardById)
	router.Post("/api/v1/cards/thehornedcard", controllers.CreateCard)
	router.Put("/api/v1/cards/thehornedcard/{id}", controllers.UpdateCard)
	router.Delete("/api/v1/cards/thehornedcard/{id}", controllers.DeleteCard)
	router.Delete("/api/v1/coffees", controllers.DeleteAllCards)

	router.Post("/api/v1/cards/groups/group", controllers.CreateGroup)

	return router
}
