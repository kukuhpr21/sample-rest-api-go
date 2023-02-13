package controllers

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
}
