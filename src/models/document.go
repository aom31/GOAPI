package models

import "gorm.io/gorm"

type Document struct {
	*gorm.Model
	Id          int    `gorm:"type:int;primary_key" json:"id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name" validate:"required"`
	Author      string `gorm:"type:varchar(50);not null" json:"author" validate:"required"`
	Description string `gorm:"type:varchar(50);not null" json:"description" validate:"required"`
}
