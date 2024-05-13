package app

import (
	"context"
	"errors"
	"fmt"
	"go-jwt/internal/pkg/logger"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
