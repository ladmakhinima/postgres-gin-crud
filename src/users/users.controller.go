package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUserAction(ctx *gin.Context) {
	var requestBody User
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request exception",
			"error":   err.Error(),
		})
		return
	}
	newUser, customErr := requestBody.SaveService()
	if customErr != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request exception",
			"error":   customErr,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created successfully",
		"data":    newUser,
	})
}

func LoadAllUsersAction(ctx *gin.Context) {
	users := FindAllService()
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func LoadUserByIdAction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println("errorrrrr", err)
	user, findErr := FindByIDService(id)
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found", "error": "not found exception"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUserByIdAction(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id is invalid",
			"error":   "bad request exception",
		})
		return
	}
	errInDelete := FindAndDeleteByIdService(id)

	if errInDelete != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
			"error":   "not found exception",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

func UpdateUserByIdAction(ctx *gin.Context) {
	var bodyRequest User
	id, errInParam := strconv.Atoi(ctx.Param("id"))
	if errInParam != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   "bad request exception",
		})
		return
	}
	err := ctx.ShouldBind(&bodyRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "bad request exception",
		})
		return
	}
	user, errInUpdate := FindAndUpdateByIdService(id, bodyRequest)
	if errInUpdate != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
			"error":   "not found exception",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "updated successfully",
	})
}
