package controllers

import (
	"errors"
	"fmt"
	usecase "jwt-demo/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	model "jwt-demo/models"
	responses "jwt-demo/responses"
	utils "jwt-demo/utils"
)

var validate = validator.New()

type UserController interface {
	Create(*gin.Context)
	Login(*gin.Context)
}

type userController struct {
	repo usecase.UserUsecase
}

var NewUserController = func(repo usecase.UserUsecase) UserController {
	return &userController{
		repo: repo,
	}
}

func (userCtl userController) Create(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}

	if err := validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}

	newUser := model.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: hashPassword,
		Phone:    user.Phone,
		Address:  user.Address,
	}

	res, err := userCtl.repo.Create(newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorInvalidData(err))
		return
	}

	c.JSON(http.StatusOK, responses.Create(*res))
}

/**
 * Use JWT to sign in user
 */
func (userCtl userController) Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}

	// if client send empty email or password
	if !(user.Email != "" && user.Password != "") {
		err := errors.New("invalid email or password")
		c.JSON(http.StatusBadRequest, responses.ErrorInvalidData(err))
		return
	}

	existedUser, err := userCtl.repo.IsExisted(user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorInvalidData(err))
		return
	}

	if !utils.CheckPasswordHash(user.Password, existedUser.Password) {
		err := errors.New("wrong email or password")
		c.JSON(http.StatusForbidden, responses.ErrorInvalidData(err))
		return
	}

	fmt.Println("success")
}
