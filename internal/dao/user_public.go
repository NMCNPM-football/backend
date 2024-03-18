package dao

import (
	"errors"
	"github.com/NMCNPM-football/backend/internal/models"
	"gorm.io/gorm"
	"strings"
)

func (u *UserDao) RegisterUserWithNewClub(user *models.User, club *models.Club) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(club).Error; err != nil {
		tx.Rollback()
		return err
	}

	user.ClubID = club.ID

	if err := tx.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate key") {
			err = tx.Unscoped().Model(&models.User{}).Where("email = ?", user.Email).Update("deleted_at", nil).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			if err = u.db.Unscoped().Where("email = ?", user.Email).Updates(&user).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (u *UserDao) RegisterUserWithExistingClub(user *models.User, club *models.Club) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	user.ClubID = club.ID

	if err := tx.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate key") {
			err = tx.Unscoped().Model(&models.User{}).Where("email = ?", user.Email).Update("deleted_at", nil).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			if err = u.db.Unscoped().Where("email = ?", user.Email).Updates(&user).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (u *UserDao) DeleteClubOwner(user *models.User, club *models.Club) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	var nextUser *models.User
	err := tx.Model(&models.User{}).Where("club_id = ?", club.ID).First(&nextUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if nextUser == nil {
		if err = tx.Delete(club).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err = tx.Model(&models.User{}).Where("id = ?", nextUser.ID).Update("position", models.CLubOwner).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (u *UserDao) DeleteClubMember(user *models.User) error {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
