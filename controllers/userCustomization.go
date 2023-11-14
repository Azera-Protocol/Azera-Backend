package controllers

import (
	"azera-backend/initializer"
	"azera-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomizationByWalletID(c *gin.Context) {
	walletID := c.Param("walletID")

	var user models.User
	// Fetch the user along with their Customization
	if err := initializer.DB.Preload("Customization").Where("wallet_id = ?", walletID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Check if Customization exists
	if user.Customization.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customization not found"})
		return
	}

	userCustomization := CustomizationResponse{
		ID:                   user.Customization.ID,
		UserID:               user.Customization.UserID,
		Username:             user.Customization.Username,
		ColorCode:            user.Customization.ColorCode,
		Image:                user.Customization.Image,
		DomainNameIdentifier: user.Customization.DomainNameIdentifier,
		DomainAddressLink:    user.Customization.DomainAddressLink,
	}

	c.JSON(http.StatusOK, userCustomization)
}

func CreateCustomizations(c *gin.Context) {
	walletID := c.Param("walletID")

	var user models.User
	if err := initializer.DB.Where("wallet_id = ?", walletID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// TODO: Handle if the user has already created a previous customization.

	var newCustomization models.Customization
	if err := c.BindJSON(&newCustomization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomization.UserID = user.ID

	// Create customizations
	if err := initializer.DB.Create(&newCustomization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customization"})
		return
	}

	userCustomization := CustomizationResponse{
		ID:                   newCustomization.ID,
		UserID:               newCustomization.UserID,
		Username:             newCustomization.Username,
		ColorCode:            newCustomization.ColorCode,
		Image:                newCustomization.Image,
		DomainNameIdentifier: newCustomization.DomainNameIdentifier,
		DomainAddressLink:    newCustomization.DomainAddressLink,
	}

	c.JSON(http.StatusOK, userCustomization)
}

func UpdateCustomizations(c *gin.Context) {
	walletID := c.Param("walletID")

	var user models.User
	if err := initializer.DB.Where("wallet_id = ?", walletID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var customization models.Customization
	if err := initializer.DB.Where("user_id = ?", user.ID).First(&customization).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customization not found"})
		return
	}

	// Bind JSON data to the customization object
	if err := c.BindJSON(&customization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the customization
	if err := initializer.DB.Save(&customization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customization"})
		return
	}

	c.JSON(http.StatusOK, customization)
}
