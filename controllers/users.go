package controllers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(ctx *gin.Context) {
	var user []*models.User
	repository.Get(&user)
	ctx.JSON(http.StatusOK, &user)
}

func Create(ctx *gin.Context) {
	user := new(models.User)
	repository.Save(&user)
	ctx.JSON(http.StatusOK, &user)
}
