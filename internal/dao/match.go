package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MatchDao struct {
	db *gorm.DB
}

func NewMatchDao(db *gorm.DB) *MatchDao {
	return &MatchDao{db}
}

func (m *MatchDao) CreateMatchCalendar(calendar *models.Matches) error {
	if err := m.db.Create(&calendar).Error; err != nil {
		return errors.Wrap(err, "c.db.Create")
	}
	return nil
}

func (m *MatchDao) GetMatchCalendarByID(id string) (*models.Matches, error) {
	var matchCalendars *models.Matches
	if err := m.db.Where("id = ?", id).Find(&matchCalendars).Error; err != nil {
		return nil, err
	}
	return matchCalendars, nil
}

func (m *MatchDao) UpdateMatchCalendar(match *models.Matches, matchID string) error {
	if err := m.db.Where("id = ?", matchID).Updates(&match).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}
func (m *MatchDao) GetAllMatchCalendars() ([]*models.Matches, error) {
	var matchCalendars []*models.Matches
	if err := m.db.Find(&matchCalendars).Error; err != nil {
		return nil, err
	}
	return matchCalendars, nil
}

func (m *MatchDao) GetAllMatchCalendarsWithStatus(status string) ([]*models.Matches, error) {
	var matchCalendars []*models.Matches
	if err := m.db.Where("status = ?", status).Find(&matchCalendars).Error; err != nil {
		return nil, err
	}
	return matchCalendars, nil
}
