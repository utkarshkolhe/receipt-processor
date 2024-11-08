package models

import (
	"utkarsh/Fetch/api/config"
  "regexp"
  "errors"
	"strconv"
	"time"
	"math"
)

// A Struct to hold a reciept
type ReceiptModel struct {
	Retailer     string      `json:"retailer" binding:"required"`
	PurchaseDate string      `json:"purchaseDate" binding:"required"`
	PurchaseTime string      `json:"purchaseTime" binding:"required"`
	Items        []ItemModel `json:"items" binding:"required"`
	Total        string      `json:"total" binding:"required"`
	Points 			 int
}

// Validate validates the ReceiptModel fields based on specified rules.
func (r *ReceiptModel) Validate() error {

	// Validate if Date matches correct format
	if _, err := time.Parse(config.DatePattern, r.PurchaseDate); err != nil {
		return errors.New(config.ErrInvalidPurchaseDateFormat)
	}

	// Validate if Time matches correct format
	if _, err := time.Parse(config.TimePattern, r.PurchaseTime); err != nil {
		return errors.New(config.ErrInvalidPurchaseTimeFormat)
	}

	// Validate if the receipt has at least min number of items
	if len(r.Items) < config.MinItemsInReceipt{
		return errors.New(config.ErrInvalidItemCount)
	}

	// Validate if total matches correct pattern
	currencyRegex := regexp.MustCompile(config.PricePattern)
	if !currencyRegex.MatchString(r.Total) {
		return errors.New(config.ErrInvalidTotalFormat)
	}

	//Validate each Item in Items
	var sum float64
	for _, item := range r.Items {
		// Validate Each Item by calling its Validation
		if err := item.Validate(); err != nil {
			return err
		}

		// Convert item price to float64 for summation
		price, err := strconv.ParseFloat(item.Price, config.FloatBase)
		if err != nil {
			return errors.New(config.ErrInvalidItemPriceFormat)
		}
		sum += price
	}

	// Convert total to float64 and compare with the sum of item prices
	total, err := strconv.ParseFloat(r.Total, config.FloatBase)
	if err != nil {
		return errors.New(config.ErrInvalidTotalFormat)
	}
	if math.Abs(total-sum) > config.Epsilon {
		return errors.New(config.ErrInvalidTotalValue)

	}


	return nil
}
