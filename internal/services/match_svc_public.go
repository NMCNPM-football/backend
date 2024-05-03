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
		return nil, fmt.Errorf("failed to get all players in club: %w", err)
	}

	response := &gen.MatchCalendarListResponse{
		Data: make([]*gen.MatchCalendarResponse_Data, 0, len(matches)),
	}

	for _, match := range matches {
		response.Data = append(response.Data, &gen.MatchCalendarResponse_Data{
			ClubOneName: match.ClubOneName,
			ClubTwoName: match.ClubTwoName,
			MatchDate:   match.MatchDate,
			MatchRound:  match.MatchRound,
			MatchStatus: match.Status,
		})
	}
	matchCount := len(matches)
	response.Message = fmt.Sprintf("There are %d matches in leauge", matchCount)
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
			MatchDate:   matchCalendar.MatchDate,
			MatchRound:  matchCalendar.MatchRound,
			MatchStatus: matchCalendar.Status,
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
			MatchDate:   matchCalendar.MatchDate,
			MatchRound:  matchCalendar.MatchRound,
			MatchStatus: matchCalendar.Status,
		},
		Message: "Match calendar fetched successfully",
	}

	return response, nil
}
