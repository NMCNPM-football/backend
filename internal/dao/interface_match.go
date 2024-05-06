package dao

import "github.com/NMCNPM-football/backend/internal/models"

type MatchDaoInterface interface {
	CreateMatchCalendar(calendar *models.Matches) error
	GetMatchCalendarByID(id string) (*models.Matches, error)
	UpdateMatchCalendar(match *models.Matches, matchID string) error
	GetAllMatchCalendars() ([]*models.Matches, error)
	GetAllMatchCalendarsWithStatus(status string) ([]*models.Matches, error)
}
