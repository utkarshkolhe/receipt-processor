package models

import (
	"utkarsh/Fetch/api/config"
  "regexp"
  "errors"
  "strconv"
)
// A Struct to hold individual item on a  reciept
type ItemModel struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

// Validate validates the fields in the ItemModel.
func (i *ItemModel) Validate() error {
  // Validate shortDescription format
  descRegex := regexp.MustCompile(config.DescriptionPattern)
	if !descRegex.MatchString(i.ShortDescription) {
		return errors.New(config.ErrInvalidShortDescriptionFormat)
	}

	// Validate price format
	priceRegex := regexp.MustCompile(config.PricePattern)
	if !priceRegex.MatchString(i.Price) {
		return errors.New(config.ErrInvalidItemPriceFormat)
	}

	// Convert price to float to ensure it's a valid number
	if _, err := strconv.ParseFloat(i.Price, config.FloatBase); err != nil {
		return errors.New(config.ErrInvalidItemPriceFormat)
	}

	return nil
}
