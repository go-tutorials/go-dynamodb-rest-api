package app

import (
	"context"

	"github.com/core-go/dynamodb"
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH = "PATCH"
)

func Route(r *mux.Router, ctx context.Context, dbConfig dynamodb.Config) error {
	app, err := NewApp(ctx, dbConfig)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.Health.Check).Methods(GET)

	userPath := "/users"
	r.HandleFunc(userPath, app.User.All).Methods(GET)
	r.HandleFunc(userPath+"/{id}", app.User.Load).Methods(GET)
	r.HandleFunc(userPath, app.User.Insert).Methods(POST)
	r.HandleFunc(userPath+"/{id}", app.User.Update).Methods(PUT)
	r.HandleFunc(userPath+"/{id}", app.User.Patch).Methods(PATCH)
	r.HandleFunc(userPath+"/{id}", app.User.Delete).Methods(DELETE)

	return nil
}
