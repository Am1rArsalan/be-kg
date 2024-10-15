package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Am1rArsalan/kelvin-green/graph"
	"github.com/Am1rArsalan/kelvin-green/repo"
	"github.com/Am1rArsalan/kelvin-green/service"
)

const defaultPort = "8080"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo := repo.NewOpenMeteoRepo()
	service := service.NewService(repo)
	graphResover := graph.NewResolver(service)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graphResover}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	http.Handle("/", corsMiddleware(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", corsMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
