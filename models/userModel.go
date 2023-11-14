package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	WalletId        string `gorm:"primaryKey"`
	TwitterUsername string `gorm:"index"`
	DiscordUsername string `gorm:"index"`
	CustomizationID uint
	Customization   Customization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
