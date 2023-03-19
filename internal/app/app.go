package app

import (
	"context"
	d "github.com/core-go/dynamodb"
	"github.com/core-go/health"

	"go-service/internal/handler"
	"go-service/internal/service"
)

type ApplicationContext struct {
	Health *health.Handler
	User   *handler.UserHandler
}

func NewApp(ctx context.Context, conf d.Config) (*ApplicationContext, error) {
	db, err := d.Connect(conf)
	if err != nil {
		return nil, err
	}

	service.CreateTableUsers(db)

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	dynamodbChecker := d.NewHealthChecker(db)
	healthHandler := health.NewHandler(dynamodbChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
