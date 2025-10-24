package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/nestjest/go-authorize-microservice/internal/app/grpc"
	"github.com/nestjest/go-authorize-microservice/internal/services/auth"
	"github.com/nestjest/go-authorize-microservice/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, TokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)

	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, TokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
