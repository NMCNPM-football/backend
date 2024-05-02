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

func (m *MatchDao) CreateMatchCalendar(calendar *models.MatchCalendar) error {
	if err := m.db.Create(&calendar).Error; err != nil {
		return errors.Wrap(err, "c.db.Create")
	}
	return nil
}

func (m *MatchDao) GetMatchCalendarByID(id string) (*models.MatchCalendar, error) {
	var calendar models.MatchCalendar
	if err := m.db.First(&calendar, id).Error; err != nil {
		return nil, errors.Wrap(err, "c.db.First")
	}
	return &calendar, nil
}

func (m *ClubDao) UpdateMatchCalendar(match *models.MatchCalendar, matchID string) error {
	if err := m.db.Where("id = ?", matchID).Updates(&match).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}
