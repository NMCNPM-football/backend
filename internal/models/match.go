package models

type Matches struct {
	ID        string `gorm:"primary_key;size:100"`
	Date      string `gorm:"size:50"`
	SeaSonID  string `gorm:"size:50"`
	Turn      string `gorm:"size:50"`
	HomeTeam  string `gorm:"size:50"`
	AwayTeam  string `gorm:"size:50"`
	MatchDate string `gorm:"size:50"`
	Location  string `gorm:"size:50"`
}

type Results struct {
	MatchID      string `gorm:"primary_key;size:100"`
	HomeTeamGoal int    `gorm:"size:50"`
	AwayTeamGoal int    `gorm:"size:50"`
	MatchStatus  string `gorm:"size:50"` // e.g. Finished, Ongoing, Postponed
	TeamWin      string `gorm:"size:50"`
	TeamLose     string `gorm:"size:50"`
	YellowCard   string `gorm:"size:50"`
	RedCard      string `gorm:"size:50"`
}

type ProgressDetail struct {
	PLayerID    string `gorm:"size:50"`
	TypeGoal    string `gorm:"size:50"`
	TimeInMatch string `gorm:"size:50"`
	Goals       string `gorm:"size:50"`
	Shots       string `gorm:"size:50"`
	ShotsOnGoal string `gorm:"size:50"`
	Passing     string `gorm:"size:50"`
}

type ProgressDetails []*ProgressDetail

type Progress struct {
	ID      string `gorm:"primary_key;size:100"`
	MatchID string `gorm:"size:50"`
	ProgressDetails
}

type MatchCalendar struct {
	Base
	ClubOneName string `gorm:"size:50"` //Home club name
	ClubTwoName string `gorm:"size:50"` //Away club name
	IntendTime  string `gorm:"size:50"`
	RealTime    string `gorm:"size:50"`
	Stadium     string `gorm:"size:50"`
	MatchRound  string `gorm:"size:50"`
	MatchTurn   string `gorm:"size:50"`
	IdClubOne   string `gorm:"size:50"`
	IdClubTwo   string `gorm:"size:50"`
	MatchStatus string `gorm:"size:50;default:'no'"`
	SeaSon      string `gorm:"size:50"`
}

type GoalType struct {
	Base
	GoalTypeName string `gorm:"size:50"`
}

type CardType struct {
	Base
	CardTypeName string `gorm:"size:50"`
}
