package app

import (
	"go-jwt/internal/config"
	delivery "go-jwt/internal/delivery/http"
	"go-jwt/internal/initializers"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/repository"
	"go-jwt/internal/server"
	"go-jwt/internal/service"
)

func (a *App) Initialize() {
	cfg, err := config.Init("configs")
	if err != nil {
		print("error")
	}

	initializers.LoadEnv()
	a.DB = initializers.ConnectToDb()
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	repos := repository.NewRepositories(a.DB)
	services := service.NewServices(service.Deps{
		Repos:  repos,
		Hasher: hasher,
	})
	handlers := delivery.NewHandler(services)
	a.Server = server.NewServer(cfg, handlers.Init(cfg))
}
