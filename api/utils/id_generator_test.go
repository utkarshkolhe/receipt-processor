package utils

import (
	"strings"
	"testing"
  "strconv"
	"github.com/stretchr/testify/assert"
)

// TestGenerateRandomString tests if generateRandomString creates strings of the correct length with valid characters.
func TestGenerateRandomString(t *testing.T) {
	// Define a list of lengths to test
	lengths := []int{1, 5, 10, 20}

	for _, length := range lengths {
		result := generateRandomString(length)

		// Check if the generated string has the correct length
		assert.Equal(t, length, len(result), "Expected length of %d but got %d", length, len(result))

		// Check if all characters in the string are from the defined CharSet
		for _, char := range result {
			assert.Contains(t, CharSet, string(char), "Character %c is not in the charset", char)
		}
	}
}

// TestGetNewID tests if GetNewID generates IDs in the correct format.
func TestGetNewID(t *testing.T) {
	// Generate a new ID
	newID := GetNewID()

	// Split the new ID by the pattern delimiter and check each segment length
	segments := strings.Split(newID, PatternDelimiter)
	expectedSegments := strings.Split(IDPattern, PatternDelimiter)

	// Check if the number of segments matches the pattern
	assert.Equal(t, len(expectedSegments), len(segments), "Expected %d segments but got %d", len(expectedSegments), len(segments))

	// Check each segment length
	for i, segment := range segments {
		expectedLength := expectedSegments[i]
		length, _ := strconv.Atoi(expectedLength)

		assert.Equal(t, length, len(segment), "Expected segment length of %d but got %d", length, len(segment))

		// Check if all characters in the segment are from the defined CharSet
		for _, char := range segment {
			assert.Contains(t, CharSet, string(char), "Character %c is not in the charset", char)
		}
	}
}
