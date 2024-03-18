package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
)

type UserDaoInterface interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	CheckExistEmail(email string) (bool, error)
	Create(user *models.User) error
	RegisterUserWithNewClub(user *models.User, club *models.Club) error
	RegisterUserWithExistingClub(user *models.User, club *models.Club) error
	DeleteClubOwner(user *models.User, club *models.Club) error
	DeleteClubMember(user *models.User) error
	Update(user *models.User) error
	ListClubMembers(companyId string) ([]*models.User, error)
}
