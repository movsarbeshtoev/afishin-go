package models


type Category struct {
	ID uint `gorm:"primariKey" json:"id"`
	Name string `json:"name" gorm:"column:name;uniqueIndex"`
	Slug string `json:"slug" gorm:"column:slug;uniqueIndex"`
}