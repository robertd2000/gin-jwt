package app

import (
	"go-jwt/internal/server"

	"gorm.io/gorm"
)

type App struct {
	Server *server.Server
	DB     *gorm.DB
}
