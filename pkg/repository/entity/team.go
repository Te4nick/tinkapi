package entity

type Team struct {
	Name               string `gorm:"primaryKey;unique"`
	SchedulingTimezone string
	Email              string
	SlackChannel       string
}
