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
)

func Init()  {
	cfg, err := config.Init("configs")
	if err != nil {
		print("error")
	}
	
	initializers.LoadEnv()
	db := initializers.ConnectToDb()
	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	fmt.Println("hasher")
	
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:                  repos,
		Hasher:                 hasher,
		})
	handlers := delivery.NewHandler(services)
	
	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
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

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}