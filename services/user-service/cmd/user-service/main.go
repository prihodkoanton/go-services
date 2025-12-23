package main

import (
	"fmt"

	"github.com/prihodkoanton/go-services/services/user-service/db"
	"github.com/prihodkoanton/go-services/services/user-service/internal/config"
	"github.com/prihodkoanton/go-services/services/user-service/internal/repository"
	"github.com/prihodkoanton/go-services/services/user-service/internal/service"
	"github.com/prihodkoanton/go-services/services/user-service/internal/transport/http"

	_ "github.com/prihodkoanton/go-services/services/user-service/internal/transport/http/docs"
	docs "github.com/prihodkoanton/go-services/services/user-service/internal/transport/http/docs"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	dbConnect, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(dbConnect)
	userService := service.NewUserService(userRepository)
	router := http.NewRouter(userService)

	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		panic(err)
	}
}
