package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"github.com/Sang-it/handy-man-api/internal-new/adapters/framework/left/graphql/generated"
	"github.com/Sang-it/handy-man-api/internal-new/adapters/framework/left/graphql/resolvers"
	"github.com/Sang-it/handy-man-api/internal-new/ports"
	"github.com/rs/cors"
)

type Adapter struct {
	db   ports.DatabasePort
	port string
}

func NewAdapter(db ports.DatabasePort, port string) *Adapter {
	return &Adapter{db: db, port: port}
}

func (httpa Adapter) Start() {
	// adapter := database.GetDatabaseAdapter()
	// db := adapter.Init()
	// httpa.db.Migrate(&handyman.HandyMan{})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", server)

	log.Printf("Running at http://localhost:%s/", environment.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+httpa.port, c.Handler(server)))
}
