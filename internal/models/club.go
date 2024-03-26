package models

type Club struct {
	Base
	NameClub    string `gorm:"size:50"`
	NameAward   string `gorm:"size:50"`
	SeaSonID    string `gorm:"size:50"`
	Shorthand   string `gorm:"size:50"`
	NameStadium string `gorm:"size:50"`
	DomainEmail string `gorm:"size:50;not null"`
	Achievement string `gorm:"size:50"`
	CreatedBy   string `gorm:"size:50"`
	UpdatedBy   string `gorm:"size:50"`
}

type Player struct {
	Base
	ClubID      string `gorm:"size:50"`
	SeaSonID    string `gorm:"size:50"`
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

type ClubPlayer struct {
	ClubID   string `gorm:"size:50"`
	PlayerID string `gorm:"size:50"`
	Role     string `gorm:"size:50"`
}

type SeaSon struct {
	SeaSonID string `gorm:"size:50;primaryKey"`
	Name     string `gorm:"size:50"`
	Country  string `gorm:"size:50"`
	Year     string `gorm:"size:50"`
}
