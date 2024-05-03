package models

type Club struct {
	Base
	OwnerBy     string `gorm:"size:50"`
	NameClub    string `gorm:"size:50"`
	NameAward   string `gorm:"size:50"`
	SeaSon      string `gorm:"size:50"`
	Shorthand   string `gorm:"size:50"`
	NameStadium string `gorm:"size:50"`
	DomainEmail string `gorm:"size:50;not null"`
	Achievement string `gorm:"size:50"`
	CreatedBy   string `gorm:"size:50"`
	UpdatedBy   string `gorm:"size:50"`
	//UpdateAt    time.Time `gorm:"autoUpdateTime"`
}

type SeaSon struct {
	SeaSonID string `gorm:"size:50;primaryKey"`
	Name     string `gorm:"size:50"`
	Country  string `gorm:"size:50"`
	Year     string `gorm:"size:50"`
}

type Coach struct {
	Base
	Name          string `gorm:"size:50"`
	NameClub      string `gorm:"size:50"`
	BirthDay      string `gorm:"size:50"`
	Country       string `gorm:"size:50"`
	Award         string `gorm:"size:50"`
	ClubID        string `gorm:"size:50"`
	Role          string `gorm:"size:50"`
	CertificateID string `gorm:"size:50"`
}

type Certificate struct {
	CertificateID string `gorm:"size:50;primaryKey"`
	Name          string `gorm:"size:50"`
	Level         string `gorm:"size:50"`
	DayReceived   string `gorm:"size:50"`
	Duration      string `gorm:"size:50"`
	IssuedBy      string `gorm:"size:50"`
	Status        string `gorm:"size:50"`
}

type Summary struct {
	Base
	TeamName       string `gorm:"size:50"`
	MatchesPlayed  int
	MatchesWon     int
	MatchesDrawn   int
	MatchesLost    int
	GoalsScored    int
	GoalsConceded  int
	YellowCards    int
	RedCards       int
	GoalDifference int
	Points         int
}
