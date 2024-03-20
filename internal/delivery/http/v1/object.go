package v1

import (
	"fmt"
	"go-jwt/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type objectCreateInput struct {
	Name        string    `json:"name" binding:"required,min=2,max=64"`
	Type        int       `json:"type" bson:"type"`
	Coords      string    `json:"coords" bson:"coords"`
	Radius      int       `json:"radius" bson:"radius"`
	Description string    `json:"description" bson:"description"`
	Color       string    `json:"color" bson:"color"`
	UserID      uuid.UUID `json:"userId" bson:"userId"`
}

func (h *Handler) initObjectsRoutes(api *gin.RouterGroup) {
	users := api.Group("/objects")
	{
		users.POST("/create", h.CreateObject)

	}
}

func (h *Handler) CreateObject(c *gin.Context) {
	var object objectCreateInput

	if err := c.BindJSON(&object); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	fmt.Println(object)

	if err := h.services.Objects.Create(c.Request.Context(), service.ObjectCreateInput{
		Name:        object.Name,
		Type:        object.Type,
		Coords:      object.Coords,
		Radius:      object.Radius,
		Description: object.Description,
		Color:       object.Color,
		UserID:      object.UserID,
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusCreated)
}
