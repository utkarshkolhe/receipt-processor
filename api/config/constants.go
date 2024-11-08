package config

// Application Constants
const(
	// Port on which to start the application
	ServerPort = "9000"
	// Name of log file
	LogFile 	 = "app.log"
	// Can touggle ConsoleLog and FileLog on and off. Writes Log to both if both are true.
	ConsoleLog = true
	FileLog    = true
)

// Field Names for API Output JSONs
const(
  FieldID     = "id"
  FieldPoints = "points"
	FieldError  = "error"
)

// Input Validation Patterns and constants
const (
  PricePattern         = `^\d+\.\d{2}$`
  DatePattern          = `2006-01-02`
	TimePattern          = `15:04`
	RetailerPattern      = `^[\w\s\-&]+$`
	DescriptionPattern   = `^[\w\s\-]+$`
  MinItemsInReceipt    = 1
)

// Floating Numbers constant
const(
  // Float difference threshold (epsilon) to consider numbers equal
	Epsilon  = 1e-9
  // Base for float conversions. can be 32 or 64
  FloatBase = 64
)
