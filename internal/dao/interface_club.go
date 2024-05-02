package dao

import (
	"github.com/NMCNPM-football/backend/internal/models"
)

type ClubDaoInterface interface {
	FindByID(id string) (*models.Club, error)
	//  FindAllClub(id string) (*models.ClubPlayer, error)
	FindByDomain(domain string) (*models.Club, error)
	FindByDomainAndSeason(domain string, season string) (*models.Club, error)
	GetClubByID(clubID string) (*models.Club, error)
	GetPLayerByID(playerID string) (*models.Player, error)
	UpdateClub(club *models.Club, clubID string) error
	UpdatePlayer(player *models.Player, playerID string) error
	GetAllClubs() ([]*models.Club, error)
	GetAllPlayersInClub(clubID string) ([]*models.Player, error)
	DeletePlayer(playerID string) error
	CreatePlayer(player *models.Player) error
}
