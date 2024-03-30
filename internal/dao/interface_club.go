package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
)

type ClubDaoInterface interface {
	FindByID(id string) (*models.Club, error)
	//  FindAllClub(id string) (*models.ClubPlayer, error)
	FindByDomain(domain string) (*models.Club, error)
	FindByDomainAndSeason(domain string, season string) (*models.Club, error)
	GetClubByName(nameClub string, name string) (*models.Club, error)
}
