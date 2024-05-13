package v1

import (
	"errors"
	"fmt"
	"go-jwt/internal/delivery/dao"
	"go-jwt/internal/domain"
	user_service "go-jwt/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.GET("/", h.getAll)
		users.PUT("/", h.updateUser)
		users.GET("/email/:email", h.getByEmail)
		users.GET("/:id", h.getByID)
		users.DELETE("/:id", h.deleteUser)

	}
}

func (h *Handler) getByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.services.Users.FindByEmail(email)

	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with email %s not found", email))
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.services.Users.FindById(userID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with id %s not found", userID))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAll(c *gin.Context) {

	users, err := h.services.Users.FindAll()

	if err != nil {
		newResponse(c, http.StatusBadRequest, "Users not found")
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) userSignUp(c *gin.Context) {
	var input dao.UserSignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	id, err := h.services.Users.SignUp(c.Request.Context(), user_service.UserSignUpInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil && !errors.Is(err, domain.ErrUserAlreadyExists) {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) userSignIn(c *gin.Context) {
	var inp dao.UserSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	t, err := h.services.Users.SignIn(c.Request.Context(), user_service.UserSignInInput{
		Email:    inp.Email,
		Password: inp.Password,
	})

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, user_service.UserSignInResponse{
		Token: t,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	var inp dao.UserUpdateInput

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	err := h.services.Users.Update(c.Request.Context(), user_service.UserUpdateInput{
		ID:    inp.ID,
		Name:  inp.Name,
		Email: inp.Email,
	})

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) deleteUser(c *gin.Context) {
	userId := c.Param("id")

	if err := h.services.Users.Delete(userId); err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Sprintf("User with id %s not found", userId))
		return
	}

	c.Status(http.StatusNoContent)
}
