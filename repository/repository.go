package repository

import (
	"context"

	"github.com/aristotekean/go-REST/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

var implementation Repository

func Close() error {
	return implementation.Close()
}

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementation.ListPosts(ctx, page)
}
