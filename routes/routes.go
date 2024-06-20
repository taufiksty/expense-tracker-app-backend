package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/expense-tracker-app-backend/controllers"
	"github.com/taufiksty/expense-tracker-app-backend/middlewares"
	"github.com/taufiksty/expense-tracker-app-backend/utils"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		utils.RespondJSON(c, http.StatusOK, "Welcome to Expense Tracker App Backend", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			utils.RespondJSON(c, http.StatusOK, "pong", nil)
		})
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		authorized := api.Group("/")
		authorized.Use(middlewares.AuthMiddleware())
		{
			authorized.DELETE("/logout", controllers.Logout)

			authorized.POST("/expenses", controllers.CreateExpense)
			authorized.GET("/expenses", controllers.GetExpenses)
			authorized.GET("/expenses/:id", controllers.GetExpenseById)
			authorized.PUT("/expenses/:id", controllers.UpdateExpenseById)
			authorized.DELETE("/expenses/:id", controllers.DeleteExpenseById)
		}
	}
}
