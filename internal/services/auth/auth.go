package auth

import (
	"context"
	"log/slog"
	"time"

	"github.com/saver89/go-grpc-auth/internal/domain/models"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, uid int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, id int) (models.App, error)
}

// New creates a new Auth service
func New(log *slog.Logger, userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		usrSaver:    userSaver,
		usrProvider: userProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
		log:         log,
	}
}

// RegisterNewUser registers a new user
func (a *Auth) RegisterNewUser(ctx context.Context, email, password string) (userID int64, err error) {
	panic("not implemented")
}

// Login logs in a user and returns a token
func (a *Auth) Login(ctx context.Context, email, password string, appID int) (token string, err error) {
	panic("not implemented")
}

// IsAdmin checks if a user is an admin
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (isAdmin bool, err error) {
	panic("not implemented")
}
