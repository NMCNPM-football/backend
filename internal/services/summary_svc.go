package services

import (
	"context"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type SummaryServiceInterface interface {
	gen.SummaryServiceServer
}

var _ SummaryServiceInterface = (*SummaryService)(nil)

type SummaryService struct {
	AbstractService
	logger     *zap.Logger
	cfg        *config.Config
	userDao    dao.UserDaoInterface
	clubDao    dao.ClubDaoInterface
	matchDao   dao.MatchDaoInterface
	summaryDao dao.SummaryDaoInterface
}

func (e *SummaryService) CreateSummary(ctx context.Context, request *gen.CreateSummaryRequest) (*gen.SuccessMessageResponse, error) {
	//TODO implement me
	return nil, nil
}

func NewSummaryService(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface, matchDao dao.MatchDaoInterface, summaryDao dao.SummaryDaoInterface) *SummaryService {
	return &SummaryService{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao, matchDao: matchDao, summaryDao: summaryDao}
}

func (e *SummaryService) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterSummaryServiceServer(s, e)
}

func (e *SummaryService) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterSummaryServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	return nil
}

//func (e *SummaryService) CreateSummary(ctx context.Context, request *gen.CreateSummaryRequest) (*gen.SuccessMessageResponse, error) {
//	// Retrieve the Results data from the database
//	results, err := e.matchDao.GetAllMatchResults()
//	if err != nil {
//		return nil, fmt.Errorf("failed to get all match done: %w", err)
//	}
//
//	// Initialize a new Summary struct
//	summary := &Summary{
//		ClubID:   results.MatchID,
//		ClubName: results.HomeTeam,
//	}
//
//	// Copy the relevant fields from Results to Summary
//	summary.YellowCard = results.YellowCardHome
//	summary.RedCard = results.RedCardHome
//
//	// Calculate the other fields based on the Results data
//	summary.MatchesPlayed = 1
//	if results.HomeTeamGoal > results.AwayTeamGoal {
//		summary.MatchesWon = 1
//	} else if results.HomeTeamGoal < results.AwayTeamGoal {
//		summary.MatchesLost = 1
//	} else {
//		summary.MatchesDraw = 1
//	}
//	summary.GoalsScored = results.HomeTeamGoal
//	summary.GoalsConceded = results.AwayTeamGoal
//	summary.GoalDifference = summary.GoalsScored - summary.GoalsConceded
//	summary.Points = summary.MatchesWon*3 + summary.MatchesDraw
//
//	// Save the new Summary to the database
//	if err := e.db.Create(&summary).Error; err != nil {
//		return nil, err
//	}
//
//	return &gen.SuccessMessageResponse{
//		Data: &gen.SuccessMessageResponseSuccessMessage{
//			Message: "Delete player successfully",
//		},
//	}, nil
//}
