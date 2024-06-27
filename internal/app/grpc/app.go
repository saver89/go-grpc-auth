package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/saver89/go-grpc-auth/internal/grpc/auth"
	"github.com/saver89/go-grpc-auth/internal/lib/logger/sl"
	"github.com/saver89/go-grpc-auth/internal/services/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	auth := auth.New(log, nil, nil, nil, 0)
	authgrpc.Register(gRPCServer, auth)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Error("failed to listen", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("address", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		log.Error("failed to serve", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	op := "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
