package dao

import "github.com/NMCNPM-football/backend/internal/models"

type MatchDaoInterface interface {
	CreateMatchCalendar(calendar *models.MatchCalendar) error
	GetMatchCalendarByID(id string) (*models.MatchCalendar, error)
	UpdateMatchCalendar(match *models.MatchCalendar, matchID string) error
	GetAllMatchCalendars() ([]*models.MatchCalendar, error)
	GetAllMatchCalendarsWithStatus(status string) ([]*models.MatchCalendar, error)
}
