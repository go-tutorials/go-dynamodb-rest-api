package app

import (
	d "github.com/core-go/dynamodb"
	"github.com/core-go/health"

	"go-service/internal/handlers"
	"go-service/internal/services"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   *handlers.UserHandler
}

func NewApp(conf d.Config) (*ApplicationContext, error) {
	db, err := d.Connect(conf)
	if err != nil {
		return nil, err
	}

	services.CreateTableUsers(db)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	dynamodbChecker := d.NewHealthChecker(db)
	healthHandler := health.NewHandler(dynamodbChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
