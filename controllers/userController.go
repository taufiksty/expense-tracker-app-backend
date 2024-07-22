package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/expense-tracker-app-backend/auth"
	"github.com/taufiksty/expense-tracker-app-backend/config"
	"github.com/taufiksty/expense-tracker-app-backend/models"
	"github.com/taufiksty/expense-tracker-app-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	config.DB.Where("email = ?", input.Email).First(&user)
	if user.ID != 0 {
		utils.RespondError(c, http.StatusBadRequest,
			fmt.Sprintf("User with email %s is already exist", input.Email))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user = models.User{
		Email:    input.Email,
		Name:     input.Name,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Registration successfully", map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid email or password")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Could not generate token")
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Login successfully", map[string]interface{}{
		"user": map[string]interface{}{
			"id":         user.ID,
			"email":      user.Email,
			"name":       user.Name,
			"updated_at": user.UpdatedAt,
		},
		"token": token,
	})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.RespondError(c, http.StatusUnauthorized, "Authorization header is required")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	blacklistToken := models.BlacklistToken{Token: tokenString}
	if err := config.DB.Create(&blacklistToken).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Could not blacklist token. Your logout was failed")
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Logout successfully", nil)
}
