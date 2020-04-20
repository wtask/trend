package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/wtask/trend/internal/graph"
)

func main() {
	log.SetPrefix("trend-graph ")
	api := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	api.SetRecoverFunc(func(_ context.Context, err interface{}) error {
		id := uuid.New().String()
		log.Println("PANIC", id, "due to:", err)
		return fmt.Errorf("internal server error [%s]", id)
	})

	http.Handle("/", playground.Handler("Trend API", "/graph"))
	http.Handle("/graph", api)

	log.Println("Server started ...")
	if err := http.ListenAndServe(":8081", nil); err != http.ErrServerClosed {
		log.Fatalln("ERROR", "Server unexpectedly stopped:", err)
	}
	// NOTE При нажатии Ctrl-C сервер чаще всего не успевает остановится
	log.Println("Server stopped, bye!")
}
