package utils

import (
	"errors"
)

// Sample texts
var sampleTexts = []string{
	"Sample text 1: This is the first sample.",
	"Sample text 2: Here’s the second sample.",
	"Sample text 3: The third sample is here.",
	"Sample text 4: And the fourth sample is this.",
}

// GetSampleTextByID returns the sample text by its ID (1-4)
// Returns an error if the ID is invalid
func GetSampleTextByID(id string) (string, error) {
	// Convert the id to an integer
	var index int
	switch id {
	case "1":
		index = 0
	case "2":
		index = 1
	case "3":
		index = 2
	case "4":
		index = 3
	default:
		return "", errors.New("invalid sample ID")
	}

	return sampleTexts[index], nil
}
