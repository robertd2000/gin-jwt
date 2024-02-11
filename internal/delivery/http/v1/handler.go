package v1

import (
	"github.com/gin-gonic/gin"
	"go-jwt/internal/service"
)

type Handler struct {
	services     *service.Services
}

func NewHandler(services *service.Services,) *Handler {
	return &Handler{
		services:     services,
		}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
	
	}
}