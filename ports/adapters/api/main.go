package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"github.com/Sang-it/handy-man-api/generated"
	"github.com/Sang-it/handy-man-api/modules/handy-man"
	"github.com/Sang-it/handy-man-api/ports/adapters/database"
	"github.com/Sang-it/handy-man-api/ports/graphql"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Start() {
	adapter := database.GetDatabaseAdapter()
	db := adapter.Init()
	db.AutoMigrate(&handyman.HandyMan{})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", server)

	log.Printf("Running at http://localhost:%s/", environment.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+environment.Get("PORT"), c.Handler(server)))

}
