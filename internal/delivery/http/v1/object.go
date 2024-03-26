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

type objectUpdateInput struct {
	ID          uuid.UUID `json:"id" bson:"id"`
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
		users.GET("/", h.getObjectsAll)
		users.GET("/:id", h.getObjectById)
		users.GET("/user/:userId", h.getObjectByUserId)
		users.POST("/", h.createObject)
		users.PUT("/", h.updateObject)
		users.DELETE("/:id", h.deleteObject)
	}
}

func (h *Handler) createObject(c *gin.Context) {
	var object objectCreateInput

	if err := c.BindJSON(&object); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

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

func (h *Handler) updateObject(c *gin.Context) {
	var object objectUpdateInput

	if err := c.BindJSON(&object); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.services.Objects.Update(c.Request.Context(), service.ObjectUpdateInput{
		ID:          object.ID,
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

func (h *Handler) getObjectsAll(c *gin.Context) {
	objects, err := h.services.Objects.FindAll()

	if err != nil {
		newResponse(c, http.StatusBadRequest, "Objects not found")

	}

	c.JSON(http.StatusOK, objects)
}

func (h *Handler) getObjectById(c *gin.Context) {
	id := c.Param("id")

	object, err := h.services.Objects.FindById(id)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("Object with id %s not found", id))
	}

	c.JSON(http.StatusOK, object)
}

func (h *Handler) getObjectByUserId(c *gin.Context) {
	userId := c.Param("userId")

	objects, err := h.services.Objects.FindByUserId(userId)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with id %s not found", userId))
	}

	c.JSON(http.StatusOK, objects)
}

func (h *Handler) deleteObject(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Objects.Delete(id)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("Object with id %s not found", id))
	}

	c.Status(http.StatusCreated)
}
