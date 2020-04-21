package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/wtask/trend/internal/graph"
)

func main() {
	logger := log.New(os.Stdout, "trend-graph ", log.LstdFlags|log.LUTC)
	api := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			logger,
			&graph.ResolverSettings{
				(&graph.ProgressionSettings{
					LengthRange: [2]int{5, 100},
				}).SetLength(15),
			},
		},
	}))
	api.SetRecoverFunc(func(_ context.Context, err interface{}) error {
		id := uuid.New().String()
		logger.Println("PANIC", id, "due to:", err)
		return fmt.Errorf("internal server error [%s]", id)
	})

	http.Handle("/", playground.Handler("Trend API", "/graph"))
	http.Handle("/graph", api)

	log.Println("Server started ...")
	if err := http.ListenAndServe(":8081", nil); err != http.ErrServerClosed {
		logger.Fatalln("ERROR", "Server unexpectedly stopped:", err)
	}
	// NOTE При нажатии Ctrl-C сервер чаще всего не успевает остановится
	logger.Println("Server stopped, bye!")
}
