package handlers

import (
	"auth-service/domain"
	"auth-service/services"
	"auth-service/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"html"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CurrentUser(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	tokenString = html.EscapeString(tokenString)

	if tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	user, err := GetUserFromToken(tokenString, uh.userService)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Token is valid", "user": user})
}

func GetUserFromToken(tokenString string, userService services.UserService) (*domain.User, error) {
	tokenString = html.EscapeString(tokenString)

	if err := utils.VerifyToken(tokenString); err != nil {
		return nil, err
	}

	claims, err := utils.ParseTokenClaims(tokenString)
	if err != nil {
		return nil, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid username in token")
	}

	user, err := userService.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uh *UserHandler) GetUserById(ctx *gin.Context) {
	userID := ctx.Param("userId")

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := uh.userService.FindUserById(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
