package controllers

import (
	"azera-backend/initializer"
	"azera-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomizationResponse struct {
	ID                   uint   `json:"id"`
	UserID               uint   `json:"user_id"`
	Username             string `json:"username"`
	ColorCode            string `json:"color_code"`
	Image                string `json:"image"`
	DomainNameIdentifier string `json:"domain_name_identifier"`
	DomainAddressLink    string `json:"domain_address_link"`
}

// A struct to get all the user relationship
type UserResponse struct {
	UserID          uint                  `json:"user_id"`
	WalletId        string                `json:"wallet_id"`
	TwitterUsername string                `json:"twitter_username"`
	DiscordUsername string                `json:"discord_username"`
	Customization   CustomizationResponse `json:"customization"`
}

// Create user response struct for the get user by wallet id
type UserFetchInfoResponse struct {
	UserID          uint   `json:"user_id"`
	WalletId        string `json:"wallet_id"`
	TwitterUsername string `json:"twitter_username"`
	DiscordUsername string `json:"discord_username"`
}

func GetUserByWalletID(c *gin.Context) {
	walletID := c.Param("walletID")

	var user models.User
	if err := initializer.DB.Preload("Customization").Where("wallet_id = ?", walletID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Map to response structs
	userResponse := UserResponse{
		UserID:          user.ID,
		WalletId:        user.WalletId,
		TwitterUsername: user.TwitterUsername,
		DiscordUsername: user.DiscordUsername,
		Customization: CustomizationResponse{
			ID:                   user.Customization.ID,
			UserID:               user.Customization.UserID,
			Username:             user.Customization.Username,
			ColorCode:            user.Customization.ColorCode,
			Image:                user.Customization.Image,
			DomainNameIdentifier: user.Customization.DomainNameIdentifier,
			DomainAddressLink:    user.Customization.DomainAddressLink,
		},
	}
	c.JSON(http.StatusOK, userResponse)
}

// Create user response struct: Defining what values I want returned once a user is created
type UserCreatedResponse struct {
	UserID          uint   `json:"user_id"`
	WalletId        string `json:"wallet_id"`
	TwitterUsername string `json:"twitter_username"`
	DiscordUsername string `json:"discord_username"`
}

// Create user function
func CreateUserDetails(c *gin.Context) {
	var newUser models.User

	// Bind JSON to newUser struct
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if WalletId already exists
	var existingUser models.User
	result := initializer.DB.Where("wallet_id = ?", newUser.WalletId).First(&existingUser)

	// Check if result found a record
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WalletId already exists"})
		return
	}

	// If WalletId does not exist, create the new user
	initializer.DB.Create(&newUser)

	// Create a UserResponse struct to send back
	userResponse := UserCreatedResponse{
		UserID:          newUser.ID,
		WalletId:        newUser.WalletId,
		TwitterUsername: newUser.TwitterUsername,
		DiscordUsername: newUser.DiscordUsername,
	}

	c.JSON(http.StatusOK, userResponse)
}

func UpdateUserSocials(c *gin.Context) {
	var user models.User
	walletID := c.Param("walletID")

	// Find user by WalletId
	if err := initializer.DB.Where("wallet_id = ?", walletID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateInfo struct {
		TwitterUsername string
		DiscordUsername string
	}
	if err := c.BindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializer.DB.Model(&user).Updates(models.User{TwitterUsername: updateInfo.TwitterUsername, DiscordUsername: updateInfo.DiscordUsername})
	// Create a UserResponse struct to send back
	userResponse := UserCreatedResponse{
		UserID:          user.ID,
		WalletId:        user.WalletId,
		TwitterUsername: user.TwitterUsername,
		DiscordUsername: user.DiscordUsername,
	}

	c.JSON(http.StatusOK, userResponse)
}
