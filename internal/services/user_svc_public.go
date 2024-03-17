package services

import (
	"context"
	"fmt"
	"github.com/NMCNPM-football/backend/common"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/gen"
	"github.com/NMCNPM-football/backend/internal/dao"
	"github.com/NMCNPM-football/backend/internal/models"
	"github.com/NMCNPM-football/backend/internal/must"
	"github.com/NMCNPM-football/backend/internal/serializers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"time"
)

type UserServicePublicInterface interface {
	gen.UserServicePublicServer
}

var _ UserServicePublicInterface = (*UserServicePublic)(nil)

type UserServicePublic struct {
	logger  *zap.Logger
	cfg     *config.Config
	userDao dao.UserDaoInterface
	clubDao dao.ClubDaoInterface
}

func NewUserServicePublic(logger *zap.Logger, cfg *config.Config, userDao dao.UserDaoInterface, clubDao dao.ClubDaoInterface) *UserServicePublic {
	return &UserServicePublic{logger, cfg, userDao, clubDao}
}

func (e *UserServicePublic) RegisterGrpcServer(s *grpc.Server) {
	gen.RegisterUserServicePublicServer(s, e)
}

func (e *UserServicePublic) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	err := gen.RegisterUserServicePublicHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	return nil
}

func (e *UserServicePublic) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (e *UserServicePublic) Register(ctx context.Context, in *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	err := common.CheckMail(in.Email)
	if err != nil {
		return nil, must.HandlerError(must.ErrInvalidEmail, e.logger)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	domainEmail := common.GetDomainEmail(in.Email)

	company, err := e.clubDao.FindByDomain(domainEmail)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	newUser := &models.User{
		Name:            in.Name,
		Email:           in.Email,
		Password:        string(hashedPassword),
		Position:        models.ClubMember,
		IsVerifiedEmail: true,
	}

	newClub := &models.Club{
		DomainEmail: domainEmail,
	}

	newRandomString := common.RandomStringNumber(10)
	if newRandomString == "" {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	newProjectName := fmt.Sprintf("p_%s", newRandomString)

	//newProject := &models.Project{
	//	Name:      newProjectName,
	//	Topic:     fmt.Sprintf("%s_%s", common.RemovePunctuation(domainEmail), newProjectName),
	//	Dataset:   fmt.Sprintf("%s_%s", common.RemovePunctuation(domainEmail), newProjectName),
	//	CompanyID: newClub.ID,
	//	CreatedBy: newUser.Email,
	//}

	if company == nil {
		newUser.Position = models.CLubOwner
		err = e.userDao.RegisterUserWithNewCompany(newUser, newClub, newProject)
		if err != nil {
			return nil, must.HandlerError(err, e.logger)
		}
	} else {
		err = e.userDao.RegisterUserWithExistingCompany(newUser, company, newProject)
		if err != nil {
			return nil, must.HandlerError(err, e.logger)
		}
	}

	registeredUser, err := e.userDao.FindByEmail(newUser.Email)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if registeredUser == nil {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	projectId, err := e.projectDao.DefaultProject(registeredUser.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if projectId == "" {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	return &gen.RegisterResponse{
		Data: &gen.RegisterResponse_Data{
			ProjectId: projectId,
			Message:   "Register successfully",
		},
	}, nil
}

func (e *UserServicePublic) RegisterWithToken(ctx context.Context, in *gen.RegisterWithTokenRequest) (*gen.RegisterResponse, error) {
	err := common.CheckMail(in.Email)
	if err != nil {
		return nil, must.HandlerError(must.ErrInvalidEmail, e.logger)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	domainEmail := common.GetDomainEmail(in.Email)

	company, err := e.companyDao.FindByDomain(domainEmail)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	invitation, err := e.projectDao.CheckInviteToken(in.Token)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if invitation == nil {
		return nil, must.HandlerError(must.ErrNoInviteFound, e.logger)
	}

	newUser := &models.User{
		Name:            in.Name,
		Email:           in.Email,
		Password:        string(hashedPassword),
		CompanyID:       company.ID,
		Position:        models.CompanyMember,
		IsVerifiedEmail: true,
	}

	newProjectMember := &models.ProjectMember{
		ProjectID: invitation.ProjectID,
		RoleID:    invitation.RoleID,
	}

	err = e.userDao.RegisterUserWithToken(newUser, newProjectMember)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	registeredUser, err := e.userDao.FindByEmail(newUser.Email)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if registeredUser == nil {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	projectId, err := e.projectDao.DefaultProject(registeredUser.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if projectId == "" {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	return &gen.RegisterResponse{
		Data: &gen.RegisterResponse_Data{
			ProjectId: projectId,
			Message:   "Register successfully",
		},
	}, nil
}

func (e *UserServicePublic) Login(ctx context.Context, in *gen.LoginRequest) (*gen.LoginResponse, error) {
	user, err := e.authenticatorByEmailPassword(in.Email, in.Password)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	data := &serializers.UserInfo{
		ID:    user.ID,
		Email: user.Email,
	}

	expire := time.Now().Add(2 * time.Hour)
	refreshExpire := time.Now().Add(24 * time.Hour)
	accessToken, refreshToken, err := must.CreateNewWithClaims(data, e.cfg.AuthenticationSecretKey, expire, refreshExpire)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	projectId, err := e.projectDao.DefaultProject(user.ID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if projectId == "" {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	return &gen.LoginResponse{
		Data: &gen.LoginResponse_Data{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ProjectId:    projectId,
		},
	}, nil
}

func (e *UserServicePublic) authenticatorByEmailPassword(email, password string) (*models.User, error) {

	user, _ := e.userDao.FindByEmail(email)
	if user != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return nil, must.ErrInvalidPassword
		}
		return user, nil
	}

	return nil, must.ErrEmailNotExists
}

func (e *UserServicePublic) DeactivateProfile(ctx context.Context, in *gen.DeactivateProfileRequest) (*gen.SuccessMessageResponse, error) {
	user, err := e.authenticatorByEmailPassword(in.Email, in.Password)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	company, err := e.companyDao.FindByID(user.CompanyID)
	if err != nil {
		return nil, must.HandlerError(err, e.logger)
	}

	if company == nil {
		return nil, must.HandlerError(must.ErrInternalServerError, e.logger)
	}

	if user.IsOwner() {
		err = e.userDao.DeleteCompanyOwner(user, company)
		if err != nil {
			return nil, must.HandlerError(err, e.logger)
		}
	} else {
		err = e.userDao.DeleteCompanyMember(user)
		if err != nil {
			return nil, must.HandlerError(err, e.logger)
		}
	}

	return &gen.SuccessMessageResponse{
		Data: &gen.SuccessMessageResponseSuccessMessage{
			Message: "Delete profile successfully",
		},
	}, nil
}
