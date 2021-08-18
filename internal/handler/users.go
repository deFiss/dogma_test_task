package handler

import (
	"dogma_test_task/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList godoc
// @Summary Get user list
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} []internal.User
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [get]
func (h *Handler) GetUserList(c *gin.Context) {
	data, err := h.services.User.GetUserList()

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, data)
}

// AddUser
// @Summary Add user
// @Tags user
// @Accept  json
// @Produce  json
// @Param input body internal.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [post]
func (h *Handler) AddUser(c *gin.Context) {
	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	id, err := h.services.AddUser(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GetUserById
// @Summary Get user by id
// @Tags user
// @ID get-string-by-int
// @Param id path string true "user id"
// @Accept  json
// @Produce  json
// @Success 200 {object} internal.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [get]
func (h *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{"invalid id"})
		return
	}

	user, err := h.services.GetUserById(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{"not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}


// UpdateUser
// @Summary Update user fields
// @Tags user
// @ID get-string-by-int
// @Param id path string true "user id"
// @Param input body internal.User true "user info"
// @Accept  json
// @Produce  json
// @Success 200 {object} internal.User
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{"invalid id"})
		return
	}

	var input internal.User

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	user, err := h.services.UpdateUser(id, input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{"not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser
// @Summary Delete user
// @Tags user
// @ID get-string-by-int
// @Param id path string true "user id"
// @Accept  json
// @Produce  json
// @Success 200 {integer} boolean true
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{"invalid id"})
		return
	}

	err = h.services.DeleteUser(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse{"not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, true)
}