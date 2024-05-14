package services

import (
	"context"
	"fmt"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	_ "github.com/NMCNPM-football/backend/internal/models"
	_ "github.com/NMCNPM-football/backend/internal/must"
	_ "github.com/NMCNPM-football/backend/internal/serializers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type MatchServicePublicInterface interface {
	gen.MatchServicePublicServer
}

var _ MatchServicePublicInterface = (*MatchServicePublic)(nil)

type MatchServicePublic struct {
	AbstractService
	logger   *zap.Logger
	cfg      *config.Config
	userDao  dao.UserDaoInterface
	clubDao  dao.ClubDaoInterface
	matchDao dao.MatchDaoInterface
}

func NewMatchServicePublic(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface, matchDao dao.MatchDaoInterface) *MatchServicePublic {
	return &MatchServicePublic{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao, matchDao: matchDao}
}

func (e *MatchServicePublic) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterMatchServicePublicServer(s, e)
}

func (e *MatchServicePublic) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterMatchServicePublicHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}
func (e *MatchServicePublic) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (e *MatchServicePublic) GetAllMatchCalendar(ctx context.Context, request *gen.EmptyRequest) (*gen.MatchCalendarListResponse, error) {
	matches, err := e.matchDao.GetAllMatchCalendars()
	if err != nil {
		return nil, fmt.Errorf("failed to get all match calendars: %w", err)
	}

	response := &gen.MatchCalendarListResponse{
		Data: make([]*gen.MatchCalendarResponse_Data, 0, len(matches)),
	}

	for _, match := range matches {
		response.Data = append(response.Data, &gen.MatchCalendarResponse_Data{
			ClubOneName: match.ClubOneName,
			ClubTwoName: match.ClubTwoName,
			IntendTime:  match.IntendTime,
			RealTime:    match.RealTime,
			MatchRound:  match.MatchRound,
			MatchTurn:   match.MatchTurn,
			MatchStatus: match.MatchStatus,
			Stadium:     match.Stadium,
			Season:      match.SeaSon,
		})
	}
	matchCount := len(matches)
	response.Message = fmt.Sprintf("There are %d matches in league", matchCount)
	return response, nil
}
func (e *MatchServicePublic) GetAllMatchCalendarsWithStatus(ctx context.Context, request *gen.StatusRequest) (*gen.MatchCalendarListResponse, error) {
	matchCalendars, err := e.matchDao.GetAllMatchCalendarsWithStatus(request.Status)
	if err != nil {
		e.logger.Error("Failed to fetch match calendars", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch match calendars: %w", err)
	}

	response := &gen.MatchCalendarListResponse{
		Data: make([]*gen.MatchCalendarResponse_Data, 0, len(matchCalendars)),
	}

	for _, matchCalendar := range matchCalendars {
		response.Data = append(response.Data, &gen.MatchCalendarResponse_Data{
			ClubOneName: matchCalendar.ClubOneName,
			ClubTwoName: matchCalendar.ClubTwoName,
			IntendTime:  matchCalendar.IntendTime,
			RealTime:    matchCalendar.RealTime,
			MatchRound:  matchCalendar.MatchRound,
			MatchStatus: matchCalendar.MatchStatus,
			Stadium:     matchCalendar.Stadium,
			Season:      matchCalendar.SeaSon,
		})
	}
	matchCount := len(matchCalendars)
	response.Message = fmt.Sprintf("There are %d matches in leauge with status '%s'", matchCount, request.Status)

	return response, nil
}

func (e *MatchServicePublic) GetMatchCalendarById(ctx context.Context, request *gen.MatchCalendarRequest) (*gen.MatchCalendarResponse, error) {
	matchCalendar, err := e.matchDao.GetMatchCalendarByID(request.Id)
	if err != nil {
		e.logger.Error("Failed to fetch match calendar", zap.Error(err))
		return nil, fmt.Errorf("failed to fetch match calendar: %w", err)
	}

	response := &gen.MatchCalendarResponse{
		Data: &gen.MatchCalendarResponse_Data{
			ClubOneName: matchCalendar.ClubOneName,
			ClubTwoName: matchCalendar.ClubTwoName,
			IntendTime:  matchCalendar.IntendTime,
			RealTime:    matchCalendar.RealTime,
			MatchTurn:   matchCalendar.MatchTurn,
			MatchRound:  matchCalendar.MatchRound,
			MatchStatus: matchCalendar.MatchStatus,
			Stadium:     matchCalendar.Stadium,
			Season:      matchCalendar.SeaSon,
		},
		Message: "Match calendar fetched successfully",
	}

	return response, nil
}

func (e *MatchServicePublic) GetMatchResultByID(ctx context.Context, request *gen.ResultScoreRequest) (*gen.ResultScoreResponse, error) {
	// Fetch the result score data using the MatchID from the request
	resultScore, err := e.matchDao.GetMatchResultByID(request.MatchId)
	if err != nil {
		return nil, fmt.Errorf("failed to get result score: %w", err)
	}

	// Create a new ResultScore
	response := &gen.ResultScoreResponse{
		Data: &gen.ResultScore{
			HomeTeamGoal:   int32(resultScore.HomeTeamGoal),
			AwayTeamGoal:   int32(resultScore.AwayTeamGoal),
			HomeTeam:       resultScore.HomeTeam,
			AwayTeam:       resultScore.AwayTeam,
			YellowCardHome: int32(resultScore.YellowCardHome),
			RedCardHome:    int32(resultScore.RedCardHome),
			YellowCardAway: int32(resultScore.YellowCardAway),
			RedCardAway:    int32(resultScore.RedCardAway),
		},
	}

	return response, nil
}

func (e *MatchServicePublic) GetAllMatchResults(ctx context.Context, request *gen.EmptyRequest) (*gen.ResultScoreListResponse, error) {
	// Fetch all result scores from the database
	resultScores, err := e.matchDao.GetAllMatchResults()
	if err != nil {
		return nil, fmt.Errorf("failed to get all match results: %w", err)
	}

	// Create a new ResultScoreListResponse
	response := &gen.ResultScoreListResponse{
		Data: make([]*gen.ResultScore, 0, len(resultScores)),
	}

	// Iterate over the result scores and add them to the response
	for _, resultScore := range resultScores {
		response.Data = append(response.Data, &gen.ResultScore{
			MatchId:        resultScore.MatchID,
			HomeTeamGoal:   int32(resultScore.HomeTeamGoal),
			AwayTeamGoal:   int32(resultScore.AwayTeamGoal),
			HomeTeam:       resultScore.HomeTeam,
			AwayTeam:       resultScore.AwayTeam,
			YellowCardHome: int32(resultScore.YellowCardHome),
			RedCardHome:    int32(resultScore.RedCardHome),
			YellowCardAway: int32(resultScore.YellowCardAway),
			RedCardAway:    int32(resultScore.RedCardAway),
		})
	}

	// Return the response
	return response, nil
}
