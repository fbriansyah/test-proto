package models

import (
	"context"

	pb "github.com/fbriansyah/test-proto/users"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Email string
	Name  string
}

type UserServer struct {
	pb.UnimplementedUserApiServer
	DB *gorm.DB
}

func (s *UserServer) GetAllUser(ctx context.Context, in *pb.UsersRequest) (*pb.UsersResponse, error) {
	users := []User{}

	result := s.DB.Find(&users)

	usersResponse := make([]*pb.User, result.RowsAffected)

	for _, user := range users {
		userResp := pb.User{
			Id:    uint64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
		usersResponse = append(usersResponse, &userResp)
	}

	return &pb.UsersResponse{Users: usersResponse}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	user := User{
		Email: in.Email,
		Name:  in.Name,
	}
	s.DB.Create(&user)

	return &pb.User{Id: uint64(user.ID), Name: user.Name, Email: user.Email}, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.User, error) {
	user := User{}

	s.DB.Where("email = ?", in.Email).Or("id = ?", in.Id).First(&user)

	return &pb.User{Id: uint64(user.ID), Name: user.Name, Email: user.Email}, nil
}
