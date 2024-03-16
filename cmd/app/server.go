package main

import (
	"context"
	"github.com/NMCNPM-football/backend/config"
	"github.com/NMCNPM-football/backend/internal/must"
	"google.golang.org/genproto/googleapis/cloud/bigquery/migration/v2"
	"google.golang.org/grpc"
	"net/http"

	"log"
	"time"
)

func main() {
	var ctx = context.TODO()
	cfg := config.ReadConfigAndArg()

	logger, sentry, err := must.NewLogger(cfg.SentryDSN, cfg.ServiceName+"-app")
	if err != nil {
		log.Fatalf("logger: %v", err)
	}

	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)

	db := must.ConnectDb(cfg.Db)
	err = migration.Migration(db)
	if err != nil {
		log.Fatalf("migration: %v", err)
	}

	opt := []grpc.ServerOption{
		//grpc
		grpc.StreamInterceptor(auth.StreamServerInterceptor(middlewareAuth.AuthMiddleware)),
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(middlewareAuth.AuthMiddleware)),
	}

	optHttpServer := []func(http.Handler) http.Handler{
		middlewareAuth.AuthHttpServerMiddleware,
	}

	//userService := services.NewUserService(logger, cfg, userDao, companyDao, projectDao)
	_, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	must.NewServer(
		ctx,
		cfg,
		opt,
		optHttpServer
		)

}
