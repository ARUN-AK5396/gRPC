package main

import (
    "context"
    "log"
    "net"

    pb "github.com/ARUN-AK5396/gRPC/proto/"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    return &pb.GetUserResponse{
        Id:    1,
        Name:  "John Doe",
        Email: "john@example.com",
    }, nil
}

func (s *server) PostUserData(ctx context.Context, req *pb.PostUserDataRequest) (*pb.PostUserDataResponse, error) {
    return &pb.PostUserDataResponse{
        Message: "Received data: " + req.Data,
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":9000")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    reflection.Register(s)
    log.Println("gRPC server running on port 9000")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
