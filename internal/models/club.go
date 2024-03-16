package models

type Club struct {
	Base
	NameClub    string `gorm:"size:50"`
	NameAward   string `gorm:"size:50"`
	Shorthand   string `gorm:"size:50"`
	NameStadium string `gorm:"size:50"`
	Achievement string `gorm:"size:50"`
	CreatedBy   string `gorm:"size:50"`
	UpdatedBy   string `gorm:"size:50"`
}

type Player struct {
	Base
	ClubID      string `gorm:"size:50"`
	SeasonID    string `gorm:"size:50"`
	TypePlayer  string `gorm:"size:50"`
	Name        string `gorm:"size:50"`
	Age         string `gorm:"size:50"`
	Height      string `gorm:"size:50"`
	Weight      string `gorm:"size:50"`
	Position    string `gorm:"size:50"`
	Nationality string `gorm:"size:50"`
	Award       string `gorm:"size:50"`
	Kit         string `gorm:"size:50"`
	Achievement string `gorm:"size:50"`
	CreatedBy   string `gorm:"size:50"`
	UpdatedBy   string `gorm:"size:50"`
}
