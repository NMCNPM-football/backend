package dao

import "github.com/NMCNPM-football/backend/internal/models"

type SummaryDaoInterface interface {
	GetSummaryByClubID(clubID string) (*models.Summary, error)
	UpdateSummary(summary *models.Summary) error
	CreateSummary(summary *models.Summary) error
	GetSummaryByClubName(clubName string) (*models.Summary, error)
	GetSummaryBySeaSon(season string) ([]*models.Summary, error)
}