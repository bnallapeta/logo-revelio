package api

import (
	"net/http"

	"github.com/bnallapeta/logo-revelio/pkg/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.LoadHTMLGlob("web/template/*")
	r.Static("/static", "./web/static")
	r.Static("/css", "./web/css")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1 := r.Group("/api/v1")
	{
		// GET
		r.GET("/game/:userID", handler.GameHandler(db))
		v1.GET("/allusers", handler.GetAllUsersHandler(db))
		v1.GET("/userscores", handler.GetUserScoresHandler(db))
		v1.GET("/toptenscores", handler.GetTopTenScoresHandler(db))

		// POST
		v1.POST("/users", handler.CreateUserHandler(db))
		v1.POST("/check-answer", handler.CheckAnswerHandler(db))
		v1.POST("/final-score", handler.UpdateFinalScoreHandler(db))
	}

	// Serve the thankyou.html file
	r.GET("/thankyou.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "thankyou.html", gin.H{})
	})
}
