package dao

import (
	"fmt"
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

func (m *MatchDao) CreateProgressScore(score *models.ProgressScore) error {
	if err := m.db.Create(&score).Error; err != nil {
		return errors.Wrap(err, "c.db.Create")
	}
	return nil
}

func (m *MatchDao) CreateResultScore(result *models.Results) error {
	if err := m.db.Create(&result).Error; err != nil {
		return err
	}
	return nil
}

func (m *MatchDao) CreateProgressCard(card *models.ProgressCard) error {
	if err := m.db.Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func (m *MatchDao) CountGoals(matchID string, clubName string) (int, error) {
	var count int
	// Write the SQL query to count the goals
	query := `SELECT COUNT(*) FROM progress_scores WHERE match_id = ? AND club_name = ?`

	// Execute the query
	err := m.db.Raw(query, matchID, clubName).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m *MatchDao) CountCard(matchID, clubName string) (int, int, error) {
	var yellowCardCount, redCardCount int

	// Query to count the number of yellow cards
	yellowCardQuery := `SELECT COUNT(*) FROM progress_cards WHERE match_id = ? AND club_name = ? AND card_type = 'LT01'`
	err := m.db.Raw(yellowCardQuery, matchID, clubName).Scan(&yellowCardCount).Error
	if err != nil {
		return 0, redCardCount, fmt.Errorf("failed to count yellow cards: %w", err)
	}

	// Query to count the number of red cards
	redCardQuery := `SELECT COUNT(*) FROM progress_cards WHERE match_id = ? AND club_name = ? AND card_type = 'LT02'`
	err = m.db.Raw(redCardQuery, matchID, clubName).Scan(&redCardCount).Error
	if err != nil {
		return yellowCardCount, 0, fmt.Errorf("failed to count red cards: %w", err)
	}

	return yellowCardCount, redCardCount, nil
}

func (m *MatchDao) GetMatchResultByID(matchID string) (*models.Results, error) {
	var matchResults *models.Results
	if err := m.db.Where("match_id = ?", matchID).First(&matchResults).Error; err != nil {
		return nil, err
	}
	return matchResults, nil
}

func (m *MatchDao) GetAllMatchResults() ([]*models.Results, error) {
	var results []*models.Results
	if err := m.db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (m *MatchDao) GetAllMatchDone() ([]*models.ProgressScore, error) {
	var matchCalendars []*models.ProgressScore
	if err := m.db.Find(&matchCalendars).Error; err != nil {
		return nil, err
	}
	return matchCalendars, nil
}
