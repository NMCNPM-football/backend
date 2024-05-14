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

func (e *SummaryService) CreateSummary(ctx context.Context, request *gen.CreateSummaryRequest) (*gen.CreateSummaryResponse, error) {
	//TODO implement me
	panic("implement me")
}
