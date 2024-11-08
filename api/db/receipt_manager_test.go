package db

import (
	"testing"
	"utkarsh/Fetch/api/models"
	"utkarsh/Fetch/api/config"
	"github.com/stretchr/testify/assert"
)

func TestAddToDatabase(t *testing.T) {
	// Create a sample receipt model with items
	receipt := models.ReceiptModel{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.ItemModel{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}

	// Call AddToDatabase function to add the receipt
	AddToDatabase("123", receipt)

	// Check if the receipt exists in the database
	_, exists := datababse["123"]
	assert.True(t, exists, "Receipt should exist in the database after being added")
}

func TestExistsInDatabase(t *testing.T) {
	// Create a sample receipt model with items
	receipt := models.ReceiptModel{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.ItemModel{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}

	// Call AddToDatabase to add the receipt
	AddToDatabase("123", receipt)

	// Check if the receipt exists
	exists := ExistsInDatabase("123")
	assert.True(t, exists, "Receipt with ID 123 should exist in the database")

	// Check for a receipt that doesn't exist
	exists = ExistsInDatabase("999")
	assert.False(t, exists, "Receipt with ID 999 should not exist in the database")
}

func TestGetReciept(t *testing.T) {
	// Create a sample receipt model with items
	receipt := models.ReceiptModel{
		Retailer:    "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.ItemModel{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}

	// Add receipt to database
	AddToDatabase("123", receipt)

	// Test for an existing receipt
	gotReceipt, err := GetReciept("123")
	assert.NoError(t, err, "There should be no error when fetching an existing receipt")
	assert.Equal(t, "Target", gotReceipt.Retailer, "The retailer should match")
	assert.Equal(t, "2022-01-01", gotReceipt.PurchaseDate, "The purchase date should match")
	assert.Equal(t, "13:01", gotReceipt.PurchaseTime, "The purchase time should match")
	assert.Equal(t, "35.35", gotReceipt.Total, "The total should match")

	// Check if items match the expected ones
	assert.Len(t, gotReceipt.Items, 5, "There should be 5 items in the receipt")
	assert.Equal(t, "Mountain Dew 12PK", gotReceipt.Items[0].ShortDescription, "The first item description should match")
	assert.Equal(t, "6.49", gotReceipt.Items[0].Price, "The price of the first item should match")

	// Test for a non-existing receipt
	gotReceipt, err = GetReciept("999")
	assert.Error(t, err, "An error should be returned when fetching a non-existing receipt")
	assert.Equal(t, config.ErrReceiptNotFound, err.Error(), "The error message should match the 'receipt not found' error")
	assert.Equal(t, models.ReceiptModel{}, gotReceipt, "The fetched receipt should be empty when not found")
}
