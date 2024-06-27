package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/saver89/go-grpc-auth/internal/app/grpc"
)

type App struct {
	GPRCSrv *grpcapp.App
}

func New(log *slog.Logger, gprcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(log, gprcPort)

	return &App{
		GPRCSrv: grpcApp,
	}
}
