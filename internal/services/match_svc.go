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

type MatchServiceInterface interface {
	gen.MatchServiceServer
}

var _ MatchServiceInterface = (*MatchService)(nil)

type MatchService struct {
	AbstractService
	logger   *zap.Logger
	cfg      *config.Config
	userDao  dao.UserDaoInterface
	clubDao  dao.ClubDaoInterface
	matchDao dao.MatchDaoInterface
}

func NewMatchService(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface, matchDao dao.MatchDaoInterface) *MatchService {
	return &MatchService{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao, matchDao: matchDao}
}

func (e *MatchService) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterMatchServiceServer(s, e)
}

func (e *MatchService) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterMatchServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *MatchService) CreateMatchCalendar(ctx context.Context, request *gen.MatchCalendar) (*gen.MatchCalendarRequest, error) {
	// Extract the MatchCalendar data from the request
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}
	if user.Position != "Admin" {
		return nil, fmt.Errorf("access denied: user is not an admin")
	}
	season := request.Season
	clubOne := request.ClubOneName
	clubTwo := request.ClubTwoName
	clubOneInfo, err := e.clubDao.FindClubByNameAndSeaSon(clubOne, season)
	if err != nil || clubOneInfo == nil {
		return nil, fmt.Errorf("failed to get club one info: %w", err)
	}
	clubTwoInfo, err := e.clubDao.FindClubByNameAndSeaSon(clubTwo, season)
	if err != nil || clubTwoInfo == nil {
		return nil, fmt.Errorf("failed to get club two info: %w", err)
	}
	// Create a new MatchCalendar model
	newMatchCalendar := &models.MatchCalendar{
		SeaSon:      request.Season,
		ClubOneName: request.ClubOneName,
		ClubTwoName: request.ClubTwoName,
		IdClubOne:   clubOneInfo.ID,
		IdClubTwo:   clubTwoInfo.ID,
		IntendTime:  request.IntendTime,
		Stadium:     clubOneInfo.Stadium,
		RealTime:    "",
		MatchRound:  request.MatchRound,
	}

	// Use the matchDao to insert the new MatchCalendar into the database
	err = e.matchDao.CreateMatchCalendar(newMatchCalendar)
	if err != nil {
		return nil, fmt.Errorf("failed to create match calendar: %w", err)
	}

	// If the insertion is successful, return a CreateMatchCalendarResponse with the ID of the new MatchCalendar
	return &gen.MatchCalendarRequest{
		Id: newMatchCalendar.ID,
	}, nil
}

func (e *MatchService) UpdateMatchCalendar(ctx context.Context, request *gen.MatchCalendar) (*gen.MatchCalendarResponse, error) {
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from context: %w", err)
	}
	if user.Position != "Admin" {
		return nil, fmt.Errorf("access denied: user is not an admin")
	}
	match, err := e.matchDao.GetMatchCalendarByID(request.Id)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}
	updateMatchCalendar := &models.MatchCalendar{
		ClubOneName: request.ClubOneName,
		ClubTwoName: request.ClubTwoName,
		IntendTime:  request.IntendTime,
		Stadium:     request.Stadium,
		RealTime:    request.RealTime,
		MatchRound:  request.MatchRound,
	}
	err = e.matchDao.UpdateMatchCalendar(updateMatchCalendar, match.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.MatchCalendarResponse{
		Data: &gen.MatchCalendarResponse_Data{
			ClubOneName: updateMatchCalendar.ClubOneName,
			ClubTwoName: updateMatchCalendar.ClubTwoName,
			IntendTime:  updateMatchCalendar.IntendTime,
			Stadium:     updateMatchCalendar.Stadium,
			RealTime:    updateMatchCalendar.RealTime,
			MatchRound:  updateMatchCalendar.MatchRound,
		},
		Message: "Match calendar updated successfully",
	}, nil
}
