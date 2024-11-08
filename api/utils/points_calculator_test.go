package utils

import (
	"testing"
	"utkarsh/Fetch/api/models"
	"github.com/stretchr/testify/assert"
)

func TestPoints1(t *testing.T) {
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
	points,err := GetPoints( receipt)
  assert.NoError(t, err, "There should be no error when generating points")
  assert.Equal(t, points, 28 ,"Expected Points %d but got %d", points, 28)
}

func TestPoints2(t *testing.T) {
	// Create a sample receipt model with items
	receipt := models.ReceiptModel{
		Retailer:    "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []models.ItemModel{
			{ShortDescription: "Gatorade", Price: "2.25"},
      {ShortDescription: "Gatorade", Price: "2.25"},
      {ShortDescription: "Gatorade", Price: "2.25"},
      {ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
	// Call AddToDatabase function to add the receipt
	points,err := GetPoints( receipt)
  assert.NoError(t, err, "There should be no error when generating points")
  assert.Equal(t, points, 109 ,"Expected Points %d but got %d", points, 109)
}
