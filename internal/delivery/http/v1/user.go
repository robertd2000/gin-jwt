package v1

import (
	"errors"
	"fmt"
	"go-jwt/internal/domain"
	"go-jwt/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.GET("/:email", h.getByEmail)
		users.GET("/", h.getAll)
		//		users.POST("/sign-in", h.userSignIn)

	}
}

func (h *Handler) getByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.services.Users.FindByEmail(email)

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("User with email %s not found", email),
		})

		return
	}

	c.JSON(200, gin.H{
		"status":     200,
		"user":       user,
		"request_at": time.Now(),
	})
}

func (h *Handler) getAll(c *gin.Context) {

	users, err := h.services.Users.FindAll()

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Users not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"status":     200,
		"users":      users,
		"request_at": time.Now(),
	})
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp userSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.services.Users.SignUp(c.Request.Context(), service.UserSignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Password: inp.Password,
	}); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusCreated)
}
