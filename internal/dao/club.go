package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ClubDao struct {
	db *gorm.DB
}

func NewClubDao(db *gorm.DB) *ClubDao {
	return &ClubDao{db}
}

func (m *ClubDao) FindClubByID(id string) (*models.Club, error) {
	var club *models.Club

	if err := m.db.Where("id = ?", id).First(&club).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}

	return club, nil
}

func (m *ClubDao) FindByDomain(domain string) (*models.Club, error) {
	var club *models.Club

	if err := m.db.Where("domain_email = ?", domain).First(&club).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return club, nil
}

func (m *ClubDao) FindByDomainAndSeason(domainEmail string, season string) (*models.Club, error) {
	var club *models.Club
	err := m.db.Where("domain_email = ? AND sea_son = ?", domainEmail, season).First(&club).Error
	if err != nil {
		return nil, err
	}
	return club, nil
}

func (m *ClubDao) GetClubByID(clubID string) (*models.Club, error) {
	var club models.Club
	if err := m.db.Where("id = ?", clubID).First(&club).Error; err != nil {
		return nil, err
	}
	return &club, nil
}

func (m *ClubDao) UpdateClub(club *models.Club, clubID string) error {
	if err := m.db.Where("id = ?", clubID).Updates(&club).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}

func (m *ClubDao) GetPLayerByID(playerID string) (*models.Player, error) {
	var player models.Player
	if err := m.db.Where("id = ?", playerID).First(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (m *ClubDao) UpdatePlayer(player *models.Player, playerID string) error {
	if err := m.db.Where("id = ?", playerID).Updates(&player).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}

func (m *ClubDao) GetAllClubs() ([]*models.Club, error) {
	var clubs []*models.Club
	if err := m.db.Find(&clubs).Error; err != nil {
		return nil, err
	}
	return clubs, nil
}

func (m *ClubDao) GetAllPlayersInClub(clubID string) ([]*models.Player, error) {
	var players []*models.Player
	if err := m.db.Where("club_id = ?", clubID).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (m *ClubDao) DeletePlayer(playerID string) error {
	if err := m.db.Where("id = ?", playerID).Delete(&models.Player{}).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Delete")
	}

	return nil
}

func (m *ClubDao) CreatePlayer(player *models.Player) error {
	if err := m.db.Create(&player).Error; err != nil {
		return errors.Wrap(err, "c.db.Create")
	}

	return nil
}

func (m *ClubDao) FindClubByNameAndSeaSon(name string, season string) (*models.Club, error) {
	var club *models.Club

	if err := m.db.Where("name_club = ? AND sea_son = ?", name, season).First(&club).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}

	return club, nil
}
