package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringArray - кастомный тип для хранения []string как JSON в БД
type StringArray []string

// Value - преобразует []string в JSON строку для сохранения в БД
func (sa StringArray) Value() (driver.Value, error) {
	if sa == nil {
		return nil, nil
	}
	return json.Marshal(sa)
}

// Scan - преобразует JSON строку из БД в []string
func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		*sa = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal StringArray value")
	}

	return json.Unmarshal(bytes, sa)
}

type Location struct {
	Venue   string `json:"venue" gorm:"column:venue"`
	Address string `json:"address" gorm:"column:address"`
	City    string `json:"city" gorm:"column:city"`
	Country string `json:"country" gorm:"column:country"`
}

type EventDate struct {
	Start     string `json:"start" gorm:"column:start_date"`
	End       string `json:"end" gorm:"column:end_date"`
	StartTime string `json:"start_time" gorm:"column:start_time"`
	EndTime   string `json:"end_time" gorm:"column:end_time"`
	Timezone  string `json:"timezone" gorm:"column:timezone"`
}

type Organizer struct {
	Name  string `json:"name" gorm:"column:organizer_name"`
	Email string `json:"email" gorm:"column:organizer_email"`
	Phone string `json:"phone" gorm:"column:organizer_phone"`
}

type EventImages struct {
	URL string `json:"url" gorm:"column:image_url"`
	Alt string `json:"alt" gorm:"column:image_alt"`
}

type EventMetadata struct {
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type Event struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `json:"title" gorm:"column:title"`
	Description     string         `json:"description" gorm:"type:text;column:description"`
	ShortDescription string        `json:"short_description" gorm:"column:short_description"`
	Location        Location       `json:"location" gorm:"embedded;embeddedPrefix:location_"`
	Date            EventDate      `json:"date" gorm:"embedded;embeddedPrefix:date_"`
	Organizer       Organizer      `json:"organizer" gorm:"embedded;embeddedPrefix:organizer_"`
	Category        string         `json:"category" gorm:"column:category"`
	Tags            StringArray      `json:"tags" gorm:"type:text;column:tags"` // JSON массив как строка
	Images          EventImages    `json:"images" gorm:"embedded;embeddedPrefix:image_"`
	Visibility      string         `json:"visibility" gorm:"column:visibility"`
	Metadata        EventMetadata  `json:"metadata" gorm:"embedded;embeddedPrefix:metadata_"`
}