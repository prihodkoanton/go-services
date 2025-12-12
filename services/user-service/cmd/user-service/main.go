package main

import (
	"strconv"

	"github.com/prihodkoanton/go-services/services/user-service/db"
	"github.com/prihodkoanton/go-services/services/user-service/internal/config"
	"github.com/prihodkoanton/go-services/services/user-service/internal/repository"
	"github.com/prihodkoanton/go-services/services/user-service/internal/service"
	"github.com/prihodkoanton/go-services/services/user-service/internal/transport/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	dbConnect, err := db.Connect(cfg)
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository(dbConnect)
	userService := service.NewUserService(userRepository)
	router := http.NewRouter(userService)

	router.Run(":", strconv.Itoa(cfg.Server.Port))
}
