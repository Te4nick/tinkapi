package entity

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Name               string `gorm:"primaryKey;unique" yaml:"name"`
	SchedulingTimezone string `yaml:"scheduling_timezone"`
	Email              string `yaml:"email"`
	SlackChannel       string `yaml:"slack_channel"`
}

type Duty struct {
	gorm.Model
	UserID string
	User   User `gorm:"references:Name"`
	TeamID string
	Team   Team `gorm:"references:Name"`
	Date   time.Time
	Role   string
}

// !!! NOT FOR DB !!! Request/Response forms

type DescDuty struct {
	Date string `yaml:"date"`
	Role string `yaml:"role"`
}

type DescUser struct {
	Name        string     `yaml:"name"`
	FullName    string     `yaml:"full_name"`
	PhoneNumber string     `yaml:"phone_number"`
	Email       string     `yaml:"email"`
	Duty        []DescDuty `yaml:"duty"`
}

type DescTeam struct {
	Name               string     `yaml:"name"`
	SchedulingTimezone string     `yaml:"scheduling_timezone"`
	Email              string     `yaml:"email"`
	SlackChannel       string     `yaml:"slack_channel"`
	Users              []DescUser `yaml:"users"`
}

type Desc struct {
	Teams []DescTeam `yaml:"teams"`
}
