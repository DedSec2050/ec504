package handlers

import (
	"eclab/utils" // Update import path to 'eclab/utils'
	"net/http"

	"github.com/atotto/clipboard" // Import the clipboard package
	"github.com/gin-gonic/gin"
)

// GetSampleText handles the API request for a specific sample text
func GetSampleText(c *gin.Context) {
	// Parse the sample id from the URL parameter
	id := c.Param("id")

	// Get the sample text based on the id
	sampleText, err := utils.GetSampleTextByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sample not found",
		})
		return
	}

	// Copy the text to the clipboard using the atotto/clipboard library
	err = clipboard.WriteAll(sampleText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to copy text to clipboard",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"text":    sampleText,
		"message": "Text successfully copied to clipboard!",
	})
}
