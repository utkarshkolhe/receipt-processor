package utils

import (
	"math/rand"
	"strconv"
	"strings"
)


// ID Generation.
const(
  // Characters to be used for ID generation. Can extend to use symbols.
  CharSet          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

  // ID format to be generated. Numbers seperated by "-". Each number will be replaced with that many characters.
  // Eg: "8-4-4-4-12" = "7fb1377b-b223-49d9-a31a-5a02701dd310"
  // Eg: "12" = "8b02701dd316"
  IDPattern        = "8-4-4-4-12"

  // Delimiter used in ID Pattern
	PatternDelimiter = "-"
)


// Function to generate a random alphanumeric string of a given length
func generateRandomString(length int) string {

	// Create a byte slice with the desired length
	randomString := make([]byte, length)

	// Fill the byte slice with random characters from the character set
	for i := range randomString {
		randomString[i] = CharSet[rand.Intn(len(CharSet))]
	}

	// Convert the byte slice to a string and return it
	return string(randomString)
}

// Function to generate a new ID based on the provided pattern. Eg: 2-3-2 = "aa-aaa-aa"
func GetNewID() string {

	// Split the pattern into segments
	segments := strings.Split(IDPattern, PatternDelimiter)
	var newIDParts []string

	// Generate a random string for each segment length
	for _, segment := range segments {
		// Convert each segment to an integer
		length, _ := strconv.Atoi(segment)

		// Generate random string of the specified length and append it to the parts
		newIDParts = append(newIDParts, generateRandomString(length))
	}

	// Join the parts with hyphens and return the result
	return strings.Join(newIDParts, PatternDelimiter)
}
