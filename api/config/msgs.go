package config


// Error messages
const (
	// Input Validation Errors
	ErrInvalidReceiptFormat          = "Invalid receipt format"
	ErrInvalidPurchaseDateFormat     = "purchaseDate must be in YYYY-MM-DD format"
	ErrInvalidPurchaseTimeFormat     = "purchaseTime must be in HH:MM 24-hour format"
	ErrInvalidTotalFormat            = "total must match currency format: xx.xx"
	ErrInvalidItemPriceFormat        = "price of each item must match currency format: xx.xx"
	ErrInvalidItemCount					     = "receipt must contain at least one item"
  ErrInvalidShortDescriptionFormat = "shortDescription of Item does not match expected pattern"
  ErrInvalidRetailerFormat         = "retailer of Reciept does not match expected pattern"
  ErrInvalidTotalValue             = "total does not match the sum of item prices"

	//Other Errors
	ErrReceiptNotFound           		 = "Receipt not found"
	ErrLogFileFailed								 = "Failed to open log file"
	ErrServerStartFailed             = "Failed to start server"
)

// Log Info Messages
const (
	// Application Start Messages
	InfoApplicationStart          	 = "Application Started"
	InfoRoutesStart          	       = "Routes Initialized"

  // GetPoints API Messages
  InfoGetPointsCalled              = "Called getPoints API"
  InfoGetPointsReturned            = "Returned getPoints API"


  //Receipt Process API Messages
  InfoPostReceiptCalled            = "Called processReceipt API"
  InfoPostReceiptReturned          = "Returned processReceipt API"
  InfoPointsCalculated             = "Points Callulated"
)
