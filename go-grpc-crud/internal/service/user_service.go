package service

import (
	"context"

	"go-grpc-crud/internal/model"
	"go-grpc-crud/internal/repository"
	"go-grpc-crud/proto/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	repo repository.UserRepository
}

// Constructor
func NewUserServiceServer(repo repository.UserRepository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}

// CreateUser
func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.UserResponse, error) {
	user := &model.User{
		Name:  req.Name,
		Email: req.Email,
	}

	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = id

	return &userpb.UserResponse{
		User: &userpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// GetUser
func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		User: &userpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// UpdateUser
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UserResponse, error) {
	user := &model.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		User: &userpb.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// DeleteUser
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteResponse, error) {
	err := s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.DeleteResponse{Success: true}, nil
}

// ListUsers
func (s *UserServiceServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var userList []*userpb.User
	for _, u := range users {
		userList = append(userList, &userpb.User{
			Id:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return &userpb.ListUsersResponse{Users: userList}, nil
}
