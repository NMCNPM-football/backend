package services

import (
	"context"
	"fmt"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/NMCNPM-football/backend/internal/must"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ClubServiceInterface interface {
	gen.ClubServiceServer
}

var _ ClubServiceInterface = (*ClubService)(nil)

type ClubService struct {
	AbstractService
	logger  *zap.Logger
	cfg     *config.Config
	userDao dao.UserDaoInterface
	clubDao dao.ClubDaoInterface
}

func (e *ClubService) GetAllPlayerProfile(ctx context.Context, request *gen.EmptyRequest) (*gen.PlayerProfileListResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}

	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	// Get all players in the club
	players, err := e.clubDao.GetAllPlayersInClub(club.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all players in club: %w", err)
	}

	// Initialize the response
	response := &gen.PlayerProfileListResponse{
		Data: make([]*gen.PLayerProfileResponse_Data, 0, len(players)),
	}

	// Convert each player to the PlayerProfile protobuf message and append it to the response
	for _, player := range players {
		response.Data = append(response.Data, &gen.PLayerProfileResponse_Data{
			ClubName:    player.ClubName,
			SeaSon:      player.SeaSon,
			TypePlayer:  player.TypePlayer,
			Name:        player.Name,
			BirthDay:    player.BirthDay,
			Height:      player.Height,
			Weight:      player.Weight,
			Position:    player.Position,
			Nationality: player.Nationality,
			Kit:         player.Kit,
			Achievement: player.Achievement,
		})

	}
	// Count the number of players in the club
	playerCount := len(players)
	response.Message = fmt.Sprintf("There are %d players in %s club", playerCount, club.NameClub)

	return response, nil

}

func NewClubService(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface) *ClubService {
	return &ClubService{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao}
}

func (e *ClubService) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterClubServiceServer(s, e)
}

func (e *ClubService) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterClubServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *ClubService) GetClubProfile(ctx context.Context, request *gen.EmptyRequest) (*gen.ClubProfileResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}
	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	// Create the response
	response := &gen.ClubProfileResponse{
		Data: &gen.ClubProfileResponse_Data{
			OwnerBy:     club.OwnerBy,
			NameClub:    club.NameClub,
			NameAward:   club.NameAward,
			SeaSon:      club.SeaSon,
			Shorthand:   club.Shorthand,
			NameStadium: club.NameStadium,
			Achievement: club.Achievement,
			UpdateBy:    club.UpdatedBy,
		},
	}

	return response, nil
}
func (e *ClubService) UpdateClub(ctx context.Context, request *gen.ClubProfileRequest) (*gen.ClubProfileResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}
	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}
	updateClub := &models.Club{
		NameClub:    request.NameClub,
		NameAward:   request.NameAward,
		NameStadium: request.NameStadium,
		SeaSon:      request.SeaSon,
		Achievement: request.Achievement,
		OwnerBy:     request.OwnerBy,
		UpdatedBy:   user.Name,
	}
	err = e.clubDao.UpdateClub(updateClub, club.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.ClubProfileResponse{
		Data: &gen.ClubProfileResponse_Data{
			OwnerBy:     updateClub.OwnerBy,
			NameClub:    updateClub.NameClub,
			NameAward:   updateClub.NameAward,
			SeaSon:      updateClub.SeaSon,
			Shorthand:   updateClub.Shorthand,
			NameStadium: updateClub.NameStadium,
			Achievement: updateClub.Achievement,
			UpdateBy:    updateClub.UpdatedBy,
		},
		Message: "Update club profile successfully",
	}, nil
}

func (e *ClubService) UpdatePlayer(ctx context.Context, request *gen.PLayerProfileRequest) (*gen.PLayerProfileResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}
	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	player, err := e.clubDao.GetPLayerByID(request.Id)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}
	updatePlayer := &models.Player{
		ClubName:    request.ClubName,
		SeaSon:      request.SeaSon,
		TypePlayer:  request.TypePlayer,
		Name:        request.Name,
		BirthDay:    request.BirthDay,
		Height:      request.Height,
		Weight:      request.Weight,
		Position:    request.Position,
		Nationality: request.Nationality,
		Kit:         request.Kit,
		Achievement: request.Achievement,
	}
	err = e.clubDao.UpdatePlayer(updatePlayer, player.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.PLayerProfileResponse{
		Data: &gen.PLayerProfileResponse_Data{
			ClubName:    updatePlayer.ClubName,
			SeaSon:      updatePlayer.SeaSon,
			TypePlayer:  updatePlayer.TypePlayer,
			Name:        updatePlayer.Name,
			BirthDay:    updatePlayer.BirthDay,
			Height:      updatePlayer.Height,
			Weight:      updatePlayer.Weight,
			Position:    updatePlayer.Position,
			Nationality: updatePlayer.Nationality,
			Kit:         updatePlayer.Kit,
			Achievement: updatePlayer.Achievement,
		},
		Message: "Update club profile successfully",
	}, nil
}

func (e *ClubService) GetPlayerProfile(ctx context.Context, request *gen.PLayerRequest) (*gen.PLayerProfileResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}
	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	player, err := e.clubDao.GetPLayerByID(request.Id)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.PLayerProfileResponse{
		Data: &gen.PLayerProfileResponse_Data{
			ClubName:    player.ClubName,
			SeaSon:      player.SeaSon,
			TypePlayer:  player.TypePlayer,
			Name:        player.Name,
			BirthDay:    player.BirthDay,
			Height:      player.Height,
			Weight:      player.Weight,
			Position:    player.Position,
			Nationality: player.Nationality,
			Kit:         player.Kit,
			Achievement: player.Achievement,
		},
		Message: "Update club profile successfully",
	}, nil

}

func (e *ClubService) GetAllClubProfile(ctx context.Context, request *gen.ClubProfileListRequest) (*gen.ClubProfileListResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Check if the user is an admin
	position := user.Position
	if position == "" {
		return nil, fmt.Errorf("failed to get user position: %w", err)
	}
	if position != "admin" {
		return nil, fmt.Errorf("user is not an admin")
	}

	clubs, err := e.clubDao.GetAllClubs()
	if err != nil {
		return nil, fmt.Errorf("failed to get all clubs: %w", err)
	}

	// Initialize the response
	response := &gen.ClubProfileListResponse{
		Data: make([]*gen.ClubProfileResponse_Data, 0, len(clubs)),
	}

	// Convert each club to the ClubProfile protobuf message and append it to the response
	for _, club := range clubs {
		response.Data = append(response.Data, &gen.ClubProfileResponse_Data{
			OwnerBy:     club.OwnerBy,
			NameClub:    club.NameClub,
			NameAward:   club.NameAward,
			SeaSon:      club.SeaSon,
			Shorthand:   club.Shorthand,
			NameStadium: club.NameStadium,
			Achievement: club.Achievement,
			UpdateBy:    club.UpdatedBy,
		})
	}

	return response, nil
}

func (e *ClubService) DeletePlayer(ctx context.Context, request *gen.PLayerRequest) (*gen.SuccessMessageResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(user.ClubID)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
	}

	// Check if the user is the owner of the club
	if user.Name != club.OwnerBy {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	// Delete the player from the club
	err = e.clubDao.DeletePlayer(request.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete player: %w", err)
	}

	return &gen.SuccessMessageResponse{
		Data: &gen.SuccessMessageResponseSuccessMessage{
			Message: "Delete player successfully",
		},
	}, nil
}
