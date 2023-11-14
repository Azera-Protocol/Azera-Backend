package models

import "gorm.io/gorm"

type Customization struct {
	gorm.Model
	UserID               uint   `gorm:"index"`
	Username             string `gorm:"index"`
	ColorCode            string `gorm:"index"`
	Image                string `gorm:"index"`
	DomainNameIdentifier string
	DomainAddressLink    string
}
