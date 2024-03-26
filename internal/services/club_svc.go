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
	//TODO implement me
	panic("implement me")
}

func (e *ClubService) UpdateClub(ctx context.Context, request *gen.ClubProfileRequest) (*gen.ClubProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}
