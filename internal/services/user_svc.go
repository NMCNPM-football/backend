package services

import (
	"context"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/NMCNPM-football/backend/internal/must"
	"github.com/NMCNPM-football/backend/internal/serializers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type UserServiceInterface interface {
	gen.UserServiceServer
}

var _ UserServiceInterface = (*UserService)(nil)

type UserService struct {
	AbstractService
	logger  *zap.Logger
	cfg     *config.Config
	userDao dao.UserDaoInterface
	clubDao dao.ClubDaoInterface
}

func NewUserService(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface) *UserService {
	return &UserService{logger: logger, cfg: cfg, userDao: userDao, clubDao: clubDao}
}

func (e *UserService) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterUserServiceServer(s, e)
}

func (e *UserService) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *UserService) RefreshToken(ctx context.Context, in *gen.EmptyRequest) (*gen.LoginResponse, error) {
	tokenInfo, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	data := &serializers.UserInfo{
		Email: tokenInfo.Email,
	}

	expire := time.Now().Add(2 * time.Hour)
	refreshExpire := time.Now().Add(24 * time.Hour)
	accessToken, refreshToken, err := must.CreateNewWithClaims(data, e.cfg.AuthenticationSecretKey, expire, refreshExpire)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	//projectId, err := e.projectDao.DefaultProject(tokenInfo.ID)
	//if err != nil {
	//	return nil, must.HandlerError(err, e.logger)
	//}
	//
	//if projectId == "" {
	//	return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	//}

	return &gen.LoginResponse{
		Data: &gen.LoginResponse_Data{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			//ProjectId:    projectId,
		},
	}, nil
}

func (e *UserService) GetProfile(ctx context.Context, in *gen.EmptyRequest) (*gen.GetProfileResponse, error) {

	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.GetProfileResponse{
		Data: &gen.GetProfileResponse_Data{
			Email:    user.Email,
			Name:     user.Name,
			Phone:    user.Phone,
			Address:  user.Address,
			ClubId:   user.ClubID,
			Position: user.Position,
		},
	}, nil
}

func (e *UserService) UpdateProfile(ctx context.Context, in *gen.UpdateProfileRequest) (*gen.SuccessMessageResponse, error) {
	user, err := e.userFromContext(ctx, e.userDao)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	newUser := &models.User{
		Email:    user.Email,
		Name:     in.Name,
		Phone:    in.Phone,
		Address:  in.Address,
		Position: in.Position,
	}

	err = e.userDao.Update(newUser)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	return &gen.SuccessMessageResponse{
		Data: &gen.SuccessMessageResponseSuccessMessage{
			Message: "Update successfully",
		},
	}, nil
}
