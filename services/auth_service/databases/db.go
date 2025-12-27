package databases

import (
	"context"

	"github.com/piyushsharma67/events_booking/services/auth_service/models"
)

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	Role         string
	Token        string
}

type Database interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
