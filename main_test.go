package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"strconv"
	"utkarsh/Fetch/api/handlers"
	"utkarsh/Fetch/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	// Set up a test router with the handlers.
	r := gin.Default()
	r.POST("/post", handlers.ProcessReciept)
	r.GET("/reciepts/:id/points", handlers.GetPoints)

	pointMap := map[string]int{
	    "Valid1": 28,
	    "Valid2": 109,
			"Valid3": 15,
			"Valid4": 31,
			"IncorrectSum":   -1, //total does not match sum of prices
			"IncorrectDate":  -1, //Invalid Date format
			"IncorrectTime":  -1, //Invalid Time format
			"IncorrectTotal": -1, //Total not in valid format
			"IncorrectPrice": -1, //Item Price not in valid format


	}
	for fileName, expectedPoints := range pointMap {
		// Open the test JSON file
		file, err := os.Open("test/"+fileName+".json")
		assert.NoError(t, err, "Failed to open test file "+fileName)
		defer file.Close()

		// Decode JSON data into a struct
		var receipt1 models.ReceiptModel
		err = json.NewDecoder(file).Decode(&receipt1)
		assert.NoError(t, err, "Failed to decode JSON data "+fileName)

		// Convert struct to JSON
		jsonData, err := json.Marshal(receipt1)
		assert.NoError(t, err, "Failed to marshal receipt to JSON "+fileName)

		// Create a POST request with the JSON data
		req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonData))
		assert.NoError(t, err, "Failed to create POST request "+fileName)
		req.Header.Set("Content-Type", "application/json")

		// Record the response
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert the status code
		if expectedPoints==-1{
				//If expectedPoints==-1 then expect 400 and move to next test case
				assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code 400 "+fileName)
				continue
		}else{
				assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200 " +fileName)
		}


		// Decode the response into an IDModel struct
		var onlineReturn models.IDModel
		err = json.Unmarshal(w.Body.Bytes(), &onlineReturn)
		assert.NoError(t, err, "Failed to unmarshal response to IDModel "+fileName)

		// Assert that an ID is returned
		assert.NotEmpty(t, onlineReturn.ID, "Expected a non-empty ID "+fileName)

		id := onlineReturn.ID

		// Create a GET request to fetch points
		req1, err := http.NewRequest("GET", "/reciepts/"+id+"/points", nil)
		assert.NoError(t, err, "Failed to create GET request "+fileName)

		// Record the response for the GET request
		w2 := httptest.NewRecorder()
		r.ServeHTTP(w2, req1)

		// Assert the status code for the GET request
		assert.Equal(t, http.StatusOK, w2.Code, "Expected status code 200 "+fileName)

		// Decode the response into a PointsModel struct
		var pointsModel models.PointsModel
		err = json.Unmarshal(w2.Body.Bytes(), &pointsModel)
		assert.NoError(t, err, "Failed to unmarshal response to PointsModel "+fileName)

		// Assert that the points are as expected
		assert.Equal(t, expectedPoints, pointsModel.Points, "Expected points to be "+strconv.Itoa(expectedPoints)+" "+fileName)
	}

}
