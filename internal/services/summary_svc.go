package services

import (
	"context"
	"fmt"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
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

func (e *SummaryService) CreateSummary(ctx context.Context, request *gen.CreateSummaryRequest) (*gen.SuccessMessageResponse, error) {

	// Fetch the match results for the club
	results, err := e.matchDao.GetAllMatchResultsBySeaSon(request.SeaSon)
	if err != nil {
		return nil, fmt.Errorf("failed to get match results for club  %w", err)
	}
	for _, result := range results {
		club, err := e.clubDao.FindClubByNameAndSeaSon(result.HomeTeam, request.SeaSon)
		if err != nil {
			return nil, fmt.Errorf("failed to find club: %w", err)
		}
		homeSummary, err := e.summaryDao.GetSummaryByClubID(club.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Initialize a new summary
				homeSummary = &models.Summary{
					ClubID:         club.ID,
					ClubName:       club.NameClub,
					MatchesPlayed:  0,
					GoalsScored:    0,
					GoalsConceded:  0,
					YellowCard:     0,
					RedCard:        0,
					MatchesWon:     0,
					MatchesDraw:    0,
					MatchesLost:    0,
					GoalDifference: 0,
					Points:         0,
					Rank:           0,
					SeaSon:         club.SeaSon,
					LogoLink:       club.LinkLogo,
				}
			}
			err = e.summaryDao.CreateSummary(homeSummary)
			if err != nil {
				return nil, fmt.Errorf("failed to create summary: %w", err)
			}
		}

		club, err = e.clubDao.FindClubByNameAndSeaSon(result.AwayTeam, request.SeaSon)
		if err != nil {
			return nil, fmt.Errorf("failed to find club: %w", err)
		}
		awaySummary, err := e.summaryDao.GetSummaryByClubID(club.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Initialize a new summary
				awaySummary = &models.Summary{
					ClubID:         club.ID,
					ClubName:       club.NameClub,
					MatchesPlayed:  0,
					GoalsScored:    0,
					GoalsConceded:  0,
					YellowCard:     0,
					RedCard:        0,
					MatchesWon:     0,
					MatchesDraw:    0,
					MatchesLost:    0,
					GoalDifference: 0,
					Points:         0,
					Rank:           0,
					SeaSon:         club.SeaSon,
					LogoLink:       club.LinkLogo,
				}
			}
			err = e.summaryDao.CreateSummary(awaySummary)
			if err != nil {
				return nil, fmt.Errorf("failed to create summary: %w", err)
			}
		}
	}
	results, err = e.matchDao.GetAllMatchResultsBySeaSon(request.SeaSon)
	if err != nil {
		return nil, fmt.Errorf("failed to get match results for club %w", err)
	}
	for _, result := range results {
		result.Status = "Finished"
		err = e.matchDao.UpdateMatch(result)
		if err != nil {
			return nil, fmt.Errorf("failed to update match result: %w", err)
		}
		// Update summary for home team
		homeSummary, err := e.summaryDao.GetSummaryByClubName(result.HomeTeam)
		if err != nil {
			return nil, fmt.Errorf("failed to get summary for club %s: %w", result.HomeTeam, err)
		}
		homeSummary.MatchesPlayed++
		homeSummary.GoalsScored += result.HomeTeamGoal
		homeSummary.GoalsConceded += result.AwayTeamGoal
		homeSummary.YellowCard += result.YellowCardHome
		homeSummary.RedCard += result.RedCardHome
		if result.HomeTeamGoal > result.AwayTeamGoal {
			homeSummary.MatchesWon++
		} else if result.HomeTeamGoal < result.AwayTeamGoal {
			homeSummary.MatchesLost++
		} else {
			homeSummary.MatchesDraw++
		}
		homeSummary.GoalDifference = homeSummary.GoalsScored - homeSummary.GoalsConceded
		homeSummary.Points = homeSummary.MatchesWon*3 + homeSummary.MatchesDraw
		err = e.summaryDao.UpdateSummary(homeSummary)
		if err != nil {
			return nil, fmt.Errorf("failed to update summary: %w", err)
		}

		// Update summary for away team
		awaySummary, err := e.summaryDao.GetSummaryByClubName(result.AwayTeam)
		if err != nil {
			return nil, fmt.Errorf("failed to get summary for club %s: %w", result.AwayTeam, err)
		}
		awaySummary.MatchesPlayed++
		awaySummary.GoalsScored += result.AwayTeamGoal
		awaySummary.GoalsConceded += result.HomeTeamGoal
		awaySummary.YellowCard += result.YellowCardAway
		awaySummary.RedCard += result.RedCardAway
		if result.AwayTeamGoal > result.HomeTeamGoal {
			awaySummary.MatchesWon++
		} else if result.AwayTeamGoal < result.HomeTeamGoal {
			awaySummary.MatchesLost++
		} else {
			awaySummary.MatchesDraw++
		}
		awaySummary.GoalDifference = awaySummary.GoalsScored - awaySummary.GoalsConceded
		awaySummary.Points = awaySummary.MatchesWon*3 + awaySummary.MatchesDraw
		err = e.summaryDao.UpdateSummary(awaySummary)
		if err != nil {
			return nil, fmt.Errorf("failed to update summary: %w", err)
		}
	}

	// Return a success message
	return &gen.SuccessMessageResponse{
		Data: &gen.SuccessMessageResponseSuccessMessage{
			Message: "Summary created successfully",
		},
	}, nil
}
