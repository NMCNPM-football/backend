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

func (c *ClubDao) FindByID(id string) (*models.Club, error) {
	var club *models.Club

	if err := c.db.Where("id = ?", id).First(&club).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}

	return club, nil
}

func (c *ClubDao) FindByDomain(domain string) (*models.Club, error) {
	var club *models.Club

	if err := c.db.Where("domain_email = ?", domain).First(&club).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return club, nil
}

func (c *ClubDao) FindByDomainAndSeason(domainEmail string, season string) (*models.Club, error) {
	var club *models.Club
	err := c.db.Where("domain_email = ? AND sea_son = ?", domainEmail, season).First(&club).Error
	if err != nil {
		return nil, err
	}
	return club, nil
}

func (c *ClubDao) GetClubByID(clubID string) (*models.Club, error) {
	var club models.Club
	if err := c.db.Where("id = ?", clubID).First(&club).Error; err != nil {
		return nil, err
	}
	return &club, nil
}

func (c *ClubDao) UpdateClub(club *models.Club, clubID string) error {
	if err := c.db.Where("id = ?", clubID).Updates(&club).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}

func (c *ClubDao) GetPLayerByID(playerID string) (*models.Player, error) {
	var player models.Player
	if err := c.db.Where("id = ?", playerID).First(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}

func (c *ClubDao) UpdatePlayer(player *models.Player, playerID string) error {
	if err := c.db.Where("id = ?", playerID).Updates(&player).Error; err != nil {
		return errors.Wrap(err, "c.db.Model.Where.Updates")
	}

	return nil
}
