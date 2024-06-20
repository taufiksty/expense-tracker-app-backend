package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/expense-tracker-app-backend/config"
	"github.com/taufiksty/expense-tracker-app-backend/models"
	"github.com/taufiksty/expense-tracker-app-backend/utils"
)

func CreateExpense(c *gin.Context) {
	var input struct {
		Amount      float64   `json:"amount" binding:"required"`
		Description string    `json:"description" binding:"required"`
		Date        time.Time `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	expense := models.Expense{
		UserID:      userID.(uint),
		Amount:      input.Amount,
		Description: input.Description,
		Date:        input.Date,
	}
	if err := config.DB.Create(&expense).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Expense created successfully", map[string]interface{}{
		"id":          expense.ID,
		"user_id":     expense.UserID,
		"amount":      expense.Amount,
		"description": expense.Description,
		"date":        expense.Date,
		"created_at":  expense.CreatedAt,
	})
}

func GetExpenses(c *gin.Context) {
	userID, _ := c.Get("userID")

	var expenses []models.Expense
	if err := config.DB.Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Expenses retrieved successfully", map[string]interface{}{
		"expenses": expenses,
	})
}

func GetExpenseById(c *gin.Context) {
	userID, _ := c.Get("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid expense ID")
	}

	var expense models.Expense
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&expense).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Expense not found")
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Expense retrieved successfully", map[string]interface{}{
		"expense": expense,
	})
}

func UpdateExpenseById(c *gin.Context) {
	userID, _ := c.Get("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid expense ID")
	}

	var expense models.Expense
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&expense).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Expense not found")
		return
	}

	var input struct {
		Amount      float64   `json:"amount"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Model(&expense).Updates(input).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Expense updated successfully", map[string]interface{}{
		"id":          expense.ID,
		"user_id":     expense.UserID,
		"amount":      expense.Amount,
		"description": expense.Description,
		"date":        expense.Date,
		"updated_at":  expense.UpdatedAt,
	})
}

func DeleteExpenseById(c *gin.Context) {
	userID, _ := c.Get("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid expense ID")
	}

	var expense models.Expense
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&expense).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Expense not found")
		return
	}

	if err := config.DB.Delete(&expense).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Expense deleted successfully", nil)
}
