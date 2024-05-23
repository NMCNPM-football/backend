package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SummaryDao struct {
	db *gorm.DB
}

func NewSummaryDao(db *gorm.DB) *SummaryDao {
	return &SummaryDao{db}
}
func (m *SummaryDao) GetSummaryByClubID(clubID string) (*models.Summary, error) {
	var summary models.Summary
	err := m.db.Where("club_id = ?", clubID).First(&summary).Error
	return &summary, err
}

func (m *SummaryDao) GetSummaryByClubName(clubName string) (*models.Summary, error) {
	var summary models.Summary
	err := m.db.Where("club_name = ?", clubName).First(&summary).Error
	return &summary, err
}

func (m *SummaryDao) UpdateSummary(summary *models.Summary) error {
	return m.db.Save(summary).Error
}

// CreateSummary Implement the CreateSummary method
func (m *SummaryDao) CreateSummary(summary *models.Summary) error {
	if err := m.db.Create(&summary).Error; err != nil {
		return errors.Wrap(err, "c.db.Create")
	}
	return nil
}

func (m *SummaryDao) GetSummaryBySeaSon(season string) ([]*models.Summary, error) {
	var summary []*models.Summary
	if err := m.db.Where("sea_son = ?", season).Find(&summary).Error; err != nil {
		return nil, err
	}
	return summary, nil
}
