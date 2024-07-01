package entities

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;type:varchar(100)"`
	Email    string `gorm:"unique;type:varchar(100)"`
	Password string `gorm:"type:varchar(255)"`
}
