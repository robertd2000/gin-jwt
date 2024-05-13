package v1

import (
	"fmt"
	"go-jwt/internal/delivery/dao"
	object_service "go-jwt/internal/service/object"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initObjectsRoutes(api *gin.RouterGroup) {
	users := api.Group("/objects")
	{
		users.GET("/", h.getAllObjects)
		users.GET("/:id", h.getObjectById)
		users.GET("/user/:userId", h.getObjectByUserId)
		users.POST("/", h.createObject)
		users.PUT("/", h.updateObject)
		users.DELETE("/:id", h.deleteObject)
	}
}

// createObject handles the HTTP POST request to create a new object.
// It expects a JSON body with the object details, and returns a StatusCreated
// response on success, or an StatusInternalServerError response on error.
//
// c *gin.Context: The Gin context object.
// Returns: None.
func (h *Handler) createObject(c *gin.Context) {
	var input dao.ObjectCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.services.Objects.Create(
		c.Request.Context(),
		object_service.ObjectCreateInput(input),
	); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// UpdateObject handles the HTTP PUT request to update an existing object.
// It expects a JSON body with the object details, and returns a StatusOK
// response on success, or an StatusInternalServerError response on error.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request.
//
// Returns: None.
func (h *Handler) updateObject(c *gin.Context) {
	var input dao.ObjectUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.services.Objects.Update(c.Request.Context(), object_service.ObjectUpdateInput(input)); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// GetAllObjects handles the HTTP GET request to retrieve all objects.
// It returns a StatusOK response with a JSON array of objects on success,
// or a StatusBadRequest response with an error message on failure.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request.
//
// Returns:
// - None.
func (h *Handler) getAllObjects(c *gin.Context) {
	objects, err := h.services.Objects.FindAll()
	if err != nil {
		newResponse(c, http.StatusBadRequest, "Failed to retrieve objects")
		return
	}

	c.JSON(http.StatusOK, objects)
}

// GetObjectById handles the HTTP GET request to retrieve a specific object.
// It takes a gin.Context as a parameter and returns a gin.Context.
//
// Parameters:
// - c: The gin.Context representing the HTTP request.
//
// Returns:
// - The modified gin.Context.
func (h *Handler) getObjectById(c *gin.Context) {
	id := c.Param("id")

	object, err := h.services.Objects.FindById(id)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("Object with id %s not found", id))
		return
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
