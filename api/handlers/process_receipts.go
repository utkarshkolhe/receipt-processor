package handlers

import (
	"net/http"
	"strconv"
	"utkarsh/Fetch/api/db"
	"utkarsh/Fetch/api/models"
	"utkarsh/Fetch/api/utils"
	"utkarsh/Fetch/api/config"
	"utkarsh/Fetch/api/logger"
	"github.com/gin-gonic/gin"
)

// Handler for end point to add reciepts to database and get back an generated id
func ProcessReciept(c *gin.Context) {

	logger.Instance.Info(config.InfoPostReceiptCalled)
	// Get the reciept from the POST Request
	var newReciept models.ReceiptModel
	if err := c.BindJSON(&newReciept); err != nil {
		// Return 400 if there was any Error with parsing the reciept.
		c.JSON(http.StatusBadRequest, gin.H{config.FieldError: err.Error()})
		logger.Instance.Error(err.Error())
		return
	}

	// Additional validation for fields
	if err := newReciept.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{config.FieldError: err.Error()})
		logger.Instance.Error(err.Error())
		return
	}

	var newId = ""
	for {
		// Get a new generated ID for the reciept
		newId = utils.GetNewID()
		// If ID is unique and does not exists in the database, break the for loop. Otherwise repeat.
		if db.ExistsInDatabase(newId) == false {
			break
		}
	}

	// Get points for the reciept
	points, err :=utils.GetPoints(newReciept)
	if err == nil {
		newReciept.Points = points
		logger.Instance.Info(config.InfoPointsCalculated+" : "+newId+" : "+strconv.Itoa(points))
	} else {
		// Return a 400 Bad Request if any error while calculating points
		c.JSON(http.StatusBadRequest, gin.H{config.FieldError: err.Error()})
		logger.Instance.Error(err.Error()+":"+newId)
	}

	// Add the reciept with the new generated id to database
	db.AddToDatabase(newId, newReciept)

	// Return the ID in JSON format
	idModel := models.IDModel{
		ID: newId,
	}
	// Returns a 200 status with id
	c.IndentedJSON(http.StatusOK, idModel)
	logger.Instance.Info(config.InfoPostReceiptReturned+" : "+newId)
}
