package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/Sang-it/handy-man-api/configs/database"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"github.com/Sang-it/handy-man-api/generated"
	handyman "github.com/Sang-it/handy-man-api/modules/handy-man"
	"github.com/Sang-it/handy-man-api/resolvers"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&handyman.HandyMan{})
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	db := database.Init()
	migrate(db)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", server)

	log.Printf("Running at http://localhost:%s/", environment.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+environment.Get("PORT"), c.Handler(server)))
}
