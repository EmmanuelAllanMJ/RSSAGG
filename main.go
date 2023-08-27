package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	// conect to the database
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("cant connect to db", err)
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of the major browsers
	}))

	// handle all requests to the /v1 path with the v1 router
	v1Route := chi.NewRouter()
	v1Route.Get("/healthz", handleRediness)
	v1Route.Get("/err", handleErr)
	v1Route.Post("/users", apiConfig.handlerCreateUser)
	v1Route.Get("/users", apiConfig.middlewareAuth(apiConfig.handlerGetUserByApiKey))

	v1Route.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlerCreateFeed))
	v1Route.Get("/feeds", apiConfig.handlerGetFeeds)
	v1Route.Post("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerCreateFeedFollow))
	v1Route.Get("/feed_follows", apiConfig.middlewareAuth(apiConfig.handlerGetFeedFollows))
	v1Route.Delete("/feed_follows/{feedFollowID}", apiConfig.middlewareAuth(apiConfig.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Route)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	log.Printf("Listening on port %s", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port: " + portString)
}
