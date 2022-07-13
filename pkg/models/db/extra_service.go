package db

import "gorm.io/gorm"

type ExtraService struct {
	gorm.Model
	Description string `gorm:"type:varchar(255);not null"`
	Price       int64
	IsActive    bool `gorm:"type:boolean;default:true;not null"`
}

/*
cihaz bakım paketi
ekran korucu
ekran koruma standart
*/
