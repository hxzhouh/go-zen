package domain

type Category struct {
	ID      string `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(32)"`
	Summary string `gorm:"type:varchar(255)"`
}
