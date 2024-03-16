package models

type User struct {
	Base
	Name      string `gorm:"size:50"`
	Email     string `gorm:"uniqueIndex:idx_email;size:50;not null"`
	Phone     string `gorm:"size:50"`
	Address   string `gorm:"size:200"`
	Password  string `gorm:"size:100"`
	CompanyID string `gorm:"size:50"`
	Position  string `gorm:"size:50"`

	IsVerifiedEmail bool   `gorm:"index:idx_is_verified_email, default:true"`
	CreatedBy       string `gorm:"size:50"`
	UpdatedBy       string `gorm:"size:50"`
	DeletedBy       string `gorm:"size:50"`
}
