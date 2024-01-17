package routes

import (
	"net/http"

	"github.com/csl991005/Atlantis/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.Mode)
	r := gin.Default()

	_ = r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "receive",
			"code":    200,
		})
	})

	r.Run(utils.Port)
}
