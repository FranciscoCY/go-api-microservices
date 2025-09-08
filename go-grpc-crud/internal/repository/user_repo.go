package repository

import (
	"context"
	"go-grpc-crud/internal/db"
	"go-grpc-crud/internal/model"
)

// UserRepository interfaz
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*model.User, error)
}

type userRepositoryGorm struct{}

func NewUserRepositoryGorm() UserRepository {
	return &userRepositoryGorm{}
}

// Create
func (r *userRepositoryGorm) Create(ctx context.Context, user *model.User) (int64, error) {
	if err := db.DB.WithContext(ctx).Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// GetByID
func (r *userRepositoryGorm) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	if err := db.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update
func (r *userRepositoryGorm) Update(ctx context.Context, user *model.User) error {
	return db.DB.WithContext(ctx).Save(user).Error
}

// Delete
func (r *userRepositoryGorm) Delete(ctx context.Context, id int64) error {
	return db.DB.WithContext(ctx).Delete(&model.User{}, id).Error
}

// List
func (r *userRepositoryGorm) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := db.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
