package services

import (
	"context"
	"fmt"

	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ClubServicePublicInterface interface {
	gen.ClubServicePublicServer
}

var _ ClubServicePublicInterface = (*ClubServicePublic)(nil)

type ClubServicePublic struct {
	AbstractService
	logger  *zap.Logger
	cfg     *config.Config
	userDao dao.UserDaoInterface
	clubDao dao.ClubDaoInterface
}

func NewClubServicePublic(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface) *ClubServicePublic {
	return &ClubServicePublic{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao}
}

func (e *ClubServicePublic) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterClubServicePublicServer(s, e)
}

func (e *ClubServicePublic) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterClubServicePublicHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *ClubServicePublic) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (e *ClubServicePublic) GetClubProfile(ctx context.Context, request *gen.ClubProfileIdRequest) (*gen.ClubProfileResponse, error) {

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(request.IdClub)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
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

func (e *ClubServicePublic) GetPlayerProfileByClub(ctx context.Context, request *gen.ClubProfileIdRequest) (*gen.PlayerProfileListResponse, error) {

	// Get the club from the context
	club, err := e.clubDao.GetClubByID(request.IdClub)
	if err != nil {
		return nil, fmt.Errorf("failed to get club by ID: %w", err)
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

func (e *ClubServicePublic) GetClubProfileListBySeaSon(ctx context.Context, request *gen.ClubProfileIdRequest) (*gen.ClubProfileListResponse, error) {

	clubs, err := e.clubDao.GetAllClubsBySeaSon(request.SeaSon)
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
			Id:          club.ID,
			OwnerBy:     club.OwnerBy,
			NameClub:    club.NameClub,
			NameAward:   club.NameAward,
			SeaSon:      club.SeaSon,
			Shorthand:   club.Shorthand,
			NameStadium: club.NameStadium,
			Achievement: club.Achievement,
			UpdateBy:    club.UpdatedBy,
			Logo:        club.LinkLogo,
		})
	}
	clubCount := len(clubs)
	response.Message = fmt.Sprintf("There are %d clubs in season 2023-2024", clubCount)
	return response, nil
}
