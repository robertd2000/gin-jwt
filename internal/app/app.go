package app

import (
	"context"
	"errors"
	"fmt"
	"go-jwt/internal/config"
	delivery "go-jwt/internal/delivery/http"
	"go-jwt/internal/initializers"
	"go-jwt/internal/pkg/hash"
	"go-jwt/internal/pkg/logger"
	"go-jwt/internal/repository"
	"go-jwt/internal/server"
	"go-jwt/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/gorm"
)

type App struct {
	Server *server.Server
	DB     *gorm.DB
}

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

func (a *App) Run() {
	go func() {
		if err := a.Server.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := a.Server.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}

func Init() {
	app := App{}

	app.Initialize()
	app.Run()
}
