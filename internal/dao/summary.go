package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
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

func (m *SummaryDao) UpdateSummary(summary *models.Summary) error {
	return m.db.Save(summary).Error
}
