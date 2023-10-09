package route

import (
	"github.com/T-sugiyama-dev/consider_buy/config"
	"github.com/T-sugiyama-dev/consider_buy/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(config.CORS())
	u := r.Group("/api/v1")
	{
		u.POST("/calcValue", controller.CalcValue)
	}

	return r
}
