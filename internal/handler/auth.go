package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marisasha/ttl-check-app/internal/models"
)

// @Summary Регистрация пользователя
// @Tags auth
// @Description Создание нового пользователя
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body models.User true "Данные пользователя"
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

}

// @Summary Аутентификация пользователя
// @Tags auth
// @Description Проверка прав пользователя
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body models.User true "Данные пользователя"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
