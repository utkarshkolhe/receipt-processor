package db

import (
	"utkarsh/Fetch/api/models"
	"errors"
	"utkarsh/Fetch/api/config"

)

// A Map used as a stand in for database
var datababse = map[string]models.ReceiptModel{}

// Add a recipt to database using ID as key
func AddToDatabase(id string, reciept models.ReceiptModel) {
	datababse[id] = reciept
}

// Check if a ID of a reciept exists in database.
func ExistsInDatabase(id string) bool {
	if _, ok := datababse[id]; ok {
		return true
	}
	return false
}

// Get a Reciept from database using its associated id. Return empty recipt if not present
func GetReciept(id string) (models.ReceiptModel,error) {
	if value, ok := datababse[id]; ok {
		return value,nil
	}
	return models.ReceiptModel{}, errors.New(config.ErrReceiptNotFound)
}
