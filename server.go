package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"github.com/Sang-it/handy-man-api/generated"
	"github.com/Sang-it/handy-man-api/modules/common"
	"github.com/Sang-it/handy-man-api/modules/handy-man"
	"github.com/Sang-it/handy-man-api/resolvers"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&handyman.HandyMan{})
}

func main() {
	db := common.Init()

	migrate(db)

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", server)

	log.Printf("Running at http://localhost:%s/", environment.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+environment.Get("PORT"), nil))
}
