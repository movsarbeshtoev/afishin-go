package models

import (
	"time"
)

const (
    
    EventStatusPending   = "pending"   
    EventStatusPublished = "published"
    EventStatusCancelled = "cancelled"
    EventStatusCompleted = "completed"
)



type EventDate struct {
	StartDate     string `json:"start_date" gorm:"column:start_date"`
	EndDate       string `json:"end_date" gorm:"column:end_date"`
	StartTime string `json:"start_time" gorm:"column:start_time"`
	EndTime   string `json:"end_time" gorm:"column:end_time"`
}

type Organizer struct {
	Name  string `json:"name" gorm:"column:organizer_name"`
	Email string `json:"email" gorm:"column:organizer_email"`
	Phone string `json:"phone" gorm:"column:organizer_phone"`
}

type EventImages struct {
	URL string `json:"url" gorm:"column:image_url"`
}

type EventMetadata struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



type Event struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `json:"title" gorm:"column:title"`
	Description     string         `json:"description" gorm:"type:text;column:description"`
	ShortDescription string        `json:"short_description" gorm:"column:short_description"`
	Address        	string       `json:"address" gorm:"column:address"`
	Date            EventDate      `json:"date" gorm:"embedded;embeddedPrefix:date_"`
	Organizer       Organizer      `json:"organizer" gorm:"embedded;embeddedPrefix:organizer_"`
	Category        string         `json:"category" gorm:"column:category"`
	Images          EventImages    `json:"images" gorm:"embedded;embeddedPrefix:image_"`
	Status      	string         `json:"status" gorm:"column:status; default:'pending'"`
	Metadata        EventMetadata  `json:"metadata" gorm:"embedded;embeddedPrefix:metadata_"`
}