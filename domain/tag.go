package domain

type Tag struct {
	ID   string `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255)"`
}
