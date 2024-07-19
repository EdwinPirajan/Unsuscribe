package domain

type Unsuscribe struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"unique"`
}
