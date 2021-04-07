package app

import (
	d "github.com/common-go/dynamodb"
	"github.com/common-go/health"
	"go-service/internal/handlers"
	"go-service/internal/services"
)


type ApplicationContext struct {
	HealthHandler *health.HealthHandler
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

	dynamodbChecker := services.NewDynamodbHealthChecker(db)
	checkers := []health.HealthChecker{dynamodbChecker}
	healthHandler := health.NewHealthHandler(checkers)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}, nil
}
