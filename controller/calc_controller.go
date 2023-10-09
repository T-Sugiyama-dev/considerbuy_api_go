package controller

import (
	"net/http"

	"github.com/T-sugiyama-dev/consider_buy/model"
	"github.com/T-sugiyama-dev/consider_buy/service"
	"github.com/gin-gonic/gin"
)

func CalcValue(c *gin.Context) {

	calcValue := model.CalcValue{}
	result := model.ResultValue{}

	if err := c.ShouldBindJSON(&calcValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	calcService := service.CalcService{}
	result, err := calcService.Calculate(&calcValue)

	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusOK, result)
}
