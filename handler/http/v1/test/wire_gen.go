// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package httphandlertest

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/configs"
	"github.com/kumin/BityDating/handler/http/v1"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos/minio"
	"github.com/kumin/BityDating/repos/mysql"
	"github.com/kumin/BityDating/repos/provider"
	"github.com/kumin/BityDating/services"
)

// Injectors from http_handler.go:

func BuildHttpHandler() (*HttpHandler, error) {
	mysqlConnector := infras.NewMysqlConnector()
	userMysqlRepo := mysql.NewUserMysqlRepo(mysqlConnector)
	client, err := infras.NewMinioClient()
	if err != nil {
		return nil, err
	}
	fileMinioRepo := minio.NewFileMinioRepo(client)
	userService := services.NewUserService(userMysqlRepo, fileMinioRepo)
	userHandler := http_handler.NewUserHandler(userService)
	matchingMysqlRepo := mysql.NewMatchingMysqlRepo(mysqlConnector)
	matchingService := services.NewMatchingService(matchingMysqlRepo)
	matchingHandler := http_handler.NewMatchingHandler(matchingService)
	authService := services.NewAuthService(userMysqlRepo)
	authHandler := http_handler.NewAuthHandler(authService)
	httpHandler := NewHttpHandler(userHandler, matchingHandler, authHandler, userService)
	return httpHandler, nil
}

// http_handler.go:

type HttpHandler struct {
	userHandler     *http_handler.UserHandler
	matchingHandler *http_handler.MatchingHandler
	authHandler     *http_handler.AuthHandler

	userService *services.UserService
}

func NewHttpHandler(
	userHandler *http_handler.UserHandler,
	matchingHandler *http_handler.MatchingHandler,
	authHandler *http_handler.AuthHandler,
	userService *services.UserService,
) *HttpHandler {
	return &HttpHandler{
		userHandler:     userHandler,
		matchingHandler: matchingHandler,
		authHandler:     authHandler,
		userService:     userService,
	}
}

var HandlerGraphSet = wire.NewSet(configs.ConfigGraphSet, services.ServiceGraphSet, provider.MysqlGraphSet, http_handler.HttpHandlerGraphSet, wire.NewSet(NewHttpHandler))
