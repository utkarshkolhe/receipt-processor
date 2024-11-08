package utils

import (
	"math"
	"strconv"
	"strings"
	"time"
  "errors"
	"unicode"
	"utkarsh/Fetch/api/models"
	"utkarsh/Fetch/api/config"
)

// You can fine tune Point Calculation by tweaking these variables
const(
  //One point for every alphanumeric character in the retailer name.
	PointAlphaNumeric       = 1

  //50 points if the total is a round dollar amount with no cents.
	PointRoundDollar        = 50
  PointNonRoundDollar     = 0

  //25 points if the total is a multiple of 0.25
  TotalDivisor            = 0.25
	PointDivisorMultiple    = 25
  PointNonDivisorMultiple = 0

  //5 points for every two items on the receipt.
  ItemSetSize          = 2
  PointItemSet         = 5

  //If the trimmed length of the item description is a multiple of 3,
  //multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
  DescriptionDivisor    = 3
  PriceMultiplier       = 0.2

  //6 points if the day in the purchase date is odd.
  PointOddPurchaseDate  = 6
  PointEvenPurchaseDate = 0

  //10 points if the time of purchase is after 2:00pm and before 4:00pm.
  StartTimeStr         = "14:00"
  EndTimeStr           = "16:00"
  PointTimeBetween		 = 10
  PointTimeNotBetween	 = 0

)

var (
    StartTime time.Time
    EndTime   time.Time
)

func init() {
    StartTime, _ = time.Parse(config.TimePattern, StartTimeStr)
    EndTime, _ = time.Parse(config.TimePattern, EndTimeStr)
}

// Calculate number of alphanumeric characters in a string
func getAlphaNumericPoints(retailerName string) int {
	count := 0

	// Iterate over each character in the string
	for _, char := range retailerName {
		// Check if the character is alphanumeric
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count*PointAlphaNumeric
}


// Calculate points if total is round dollar amount with no cents
func getRoundDollarPoints(total float64) int {
  if total == math.Trunc(total) {
		return PointRoundDollar
	}
	return PointNonRoundDollar
}

// Calculate points if total is round multiple of a divisor(0.25)
func getRoundPartMultiple(total float64) int {
	remainder := math.Mod(total, TotalDivisor)
	if math.Abs(remainder) < config.Epsilon {
		return PointDivisorMultiple
	}
	return PointNonDivisorMultiple
}

// Calculate points for every ItemSetSize(2) Items in Reciepts
func getItemSetPoints(numItems int) int {
	return (PointItemSet * (numItems / ItemSetSize))
}

// Calculate Points for all items whose trimmed description is divisible by DescriptionDivisor(3)
func getItemDescriptionPoints(items []models.ItemModel) (int,error) {
  points:=0
  for _, item := range items {
    // Trim the item short description
    trimmed := strings.TrimSpace(item.ShortDescription)
    // Add points if trimmed desc length is divisible by 3
    if len(trimmed)%DescriptionDivisor == 0 {
      price, err := strconv.ParseFloat(item.Price, config.FloatBase)
      if err != nil {
        // Handle the error if the conversion fails
        return points,errors.New(config.ErrInvalidItemPriceFormat)
      }
      // Multiply price by associated multiplier to get points
      points += int(math.Ceil(price * PriceMultiplier))
    }
  }
  return points,nil
}

func getOddEvenDayPoints(date time.Time) int{
  // Extract the day component from the parsed date
	day := date.Day()
  if day%2 == 1 {
		return PointOddPurchaseDate
	} else{
    return PointEvenPurchaseDate
  }
}

func getTimeBetweenRangePoints(time1 time.Time) int{
  // Extract the day component from the parsed date
  if time1.After(StartTime) && time1.Before(EndTime) {
		return PointTimeBetween
	} else{
    return PointTimeNotBetween
  }
}



// Calculate and return points for a ID of reciept. Returns in string format.
func GetPoints(receipt models.ReceiptModel) (int,error) {

	var points = 0

	// Add number of alpanumeric characters as points
	points += getAlphaNumericPoints(receipt.Retailer)

  // Convert total points to float
	total, err := strconv.ParseFloat(receipt.Total, config.FloatBase)
	if err != nil {
		// Handle the error if the conversion fails
		return points,errors.New(config.ErrInvalidTotalValue)
	}

  //Add points for round Dollar
  points += getRoundDollarPoints(total)

  //Add
  points += getRoundPartMultiple(total)

  //Add 5 points for every two pairs in items
	points += getItemSetPoints(len(receipt.Items))

	// For every item trimmed short description add correspoinding poinst
	itemDescPoints,err := getItemDescriptionPoints(receipt.Items)
  if err!= nil{
    return points,err
  }
  points += itemDescPoints
	// Parse the date string into a time.Time object
	date, err := time.Parse(config.DatePattern, receipt.PurchaseDate)
	if err != nil {
		return points, errors.New(config.ErrInvalidPurchaseDateFormat)
	}

	points += getOddEvenDayPoints(date)

	// Parse the Purchase time into a time object
	purchaseTime, err := time.Parse(config.TimePattern, receipt.PurchaseTime)
	if err != nil {
		return points, errors.New(config.ErrInvalidPurchaseTimeFormat)
	}
	// Add 10 poinst if purchase time between 2pm and 4pm
	points += getTimeBetweenRangePoints(purchaseTime)

	return points,nil

}
