package controllers

import (
	"github.com/alganbr/kedai-usersvc-client/client"
	"github.com/alganbr/kedai-usersvc-client/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IUsersController interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
}

type UsersController struct {
	svc client.IUserSvcClient
}

func NewUsersController(svc client.IUserSvcClient) IUsersController {
	return &UsersController{
		svc: svc,
	}
}

// Get godoc
// @Security ApiKeyAuth
// @Description  Get user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "User ID"
// @Success      200  {object}  models.User
// @Router       /users/{id} [get]
func (ctrl *UsersController) Get(c *gin.Context) {
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, getErr := ctrl.svc.User().Get(id)
	if getErr != nil {
		c.AbortWithStatusJSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Create godoc
// @Security ApiKeyAuth
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      201      {object}  models.User
// @Router       /users [post]
func (ctrl *UsersController) Create(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	result, saveErr := ctrl.svc.User().Create(&rq)
	if saveErr != nil {
		c.AbortWithStatusJSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Update godoc
// @Security ApiKeyAuth
// @Description  Update existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "User ID"
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      200      {object}  models.User
// @Router       /users/{id} [put]
func (ctrl *UsersController) Update(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, updateErr := ctrl.svc.User().Update(id, &rq)
	if updateErr != nil {
		c.AbortWithStatusJSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Patch godoc
// @Security ApiKeyAuth
// @Description  Update existing user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request  body      models.UserRq  true  "User Request"
// @Success      200      {object}  models.User
// @Router       /users/{id} [patch]
func (ctrl *UsersController) Patch(c *gin.Context) {
	var rq models.UserRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, patchErr := ctrl.svc.User().Patch(id, &rq)
	if patchErr != nil {
		c.AbortWithStatusJSON(patchErr.Code, patchErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
