package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/saver89/go-grpc-auth/internal/app/grpc"
	"github.com/saver89/go-grpc-auth/internal/services/auth"
	"github.com/saver89/go-grpc-auth/internal/storage/sqlite"
)

type App struct {
	GPRCSrv *grpcapp.App
}

func New(log *slog.Logger, gprcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	// TODO: init auth service

	grpcApp := grpcapp.New(log, authService, gprcPort)

	return &App{
		GPRCSrv: grpcApp,
	}
}
