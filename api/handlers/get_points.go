package handlers

import (
	"net/http"
	"utkarsh/Fetch/api/db"
	"utkarsh/Fetch/api/models"
	"utkarsh/Fetch/api/config"
	"utkarsh/Fetch/api/logger"
	"strconv"
	"github.com/gin-gonic/gin"
)

// Handler for point request end point. Get points for a reciept associated with an id
func GetPoints(c *gin.Context) {
	// Get id from get request
	id := c.Param(config.FieldID)
	logger.Instance.Info(config.InfoGetPointsCalled+":"+id)

	// Get points for the associated reciept and convert it to int if possible.
	reciept,err := db.GetReciept(id)
	if err!=nil {
		// Return a 400 Bad Request response if reciept not present in Database
		c.JSON(http.StatusBadRequest, gin.H{config.FieldError: err.Error()})
		logger.Instance.Error(err.Error()+":"+id)
		return
	}else{
		pointsModel := models.PointsModel{
			Points: reciept.Points,
		}
		// Returns a 200 Status with pointsModel
		c.IndentedJSON(http.StatusOK, pointsModel)
		logger.Instance.Info(config.InfoGetPointsReturned+" : "+id+" : "+strconv.Itoa(reciept.Points))
	}

}
