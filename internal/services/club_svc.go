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

func NewClubService(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface) *ClubService {
	return &ClubService{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao}
}

func (e *ClubService) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterClubServiceServer(s, e)
}

func (e *ClubService) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *ClubService) GetClubProfile(ctx context.Context, request *gen.EmptyRequest) (*gen.ClubProfileResponse, error) {
	// Get the user from the context
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, err
	}

	// Check if the user is the owner of the club
	if !user.IsOwner() {
		return nil, fmt.Errorf("user is not the owner of the club")
	}

	// Get the club profile from the database
	club, err := e.clubDao.GetClubByName(user.Club, user.Name)
	if err != nil {
		return nil, err
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
			DomainEmail: club.DomainEmail,
			Achievement: club.Achievement,
			UpdateBy:    club.UpdatedBy,
		},
	}

	return response, nil
}
func (e *ClubService) UpdateClub(ctx context.Context, request *gen.ClubProfileRequest) (*gen.ClubProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}
