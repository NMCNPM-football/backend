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

func (u *ClubDao) FindByID(id string) (*models.Club, error) {
	var company *models.Club

	if err := u.db.Where("id = ?", id).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}

	return company, nil
}

func (u *ClubDao) FindByDomain(domain string) (*models.Club, error) {
	var company *models.Club

	if err := u.db.Where("domain_email = ?", domain).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "u.db.Where.First")
	}
	return company, nil
}
