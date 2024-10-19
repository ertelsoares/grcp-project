package server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/pb"
	"net"
	"sync"
)

type User struct {
	ID       string
	UserName string
	Password string
}

type UserService struct {
	pb.UnimplementedUserServer

	users map[string]*User
	mu    sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}
func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, NewUserService())
	reflection.Register(s)
	fmt.Println("gRPC server is running on port 50051...")
	err = s.Serve(listen)

	if err != nil {
		panic(err)
	}
}
func (us *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := &User{
		ID:       req.Id,
		UserName: req.Username,
		Password: req.Password,
	}

	us.users[user.ID] = user

	return &pb.AddUserResponse{
		Id:       user.ID,
		Username: user.UserName,
		Password: user.Password,
	}, nil
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	user, ok := us.users[req.Id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{
		Id:       user.ID,
		Username: user.UserName,
		Password: user.Password,
	}, nil
}