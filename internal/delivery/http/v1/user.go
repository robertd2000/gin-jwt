package v1

import (
	"errors"
	"fmt"
	"go-jwt/internal/domain"
	"go-jwt/internal/service"
	"net/http"

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
		users.GET("/", h.getAll)
		users.GET("/email/:email", h.getByEmail)
		users.GET("/:id", h.getById)
		users.POST("/sign-in", h.userSignIn)

	}
}

func (h *Handler) getByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.services.Users.FindByEmail(email)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with email %s not found", email))
	}

	c.JSON(200, user)
}

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.Users.FindById(id)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with id %s not found", id))
	}

	c.JSON(200, user)
}

func (h *Handler) getAll(c *gin.Context) {

	users, err := h.services.Users.FindAll()

	if err != nil {
		newResponse(c, http.StatusBadRequest, "Users not found")
	}

	c.JSON(200, users)
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

func (h *Handler) userSignIn(c *gin.Context) {
	var inp userSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	t, err := h.services.Users.SignIn(c.Request.Context(), service.UserSignInInput{
		Email:    inp.Email,
		Password: inp.Password,
	})

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, service.UserSignInResponse{
		Token: t,
	})

}
