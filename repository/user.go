package repository

type User struct {
	Account  string `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Gender   int    `gorm:"DEFAULT:'1';not null"`
	EMail    string `gorm:"null"`
}
