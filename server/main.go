package main

import (
	"flag"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "user/pb"
    "log"
    "net"
)

type userService struct{}

func (e *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return &pb.User{ Name: "watanabeJunna" }, nil
}

func main() {
	var port string

	flag.StringVar(&port, "port", "9090", "")
	flag.Parse()

	listen, err := net.Listen("tcp", ":" + port)

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})
	reflection.Register(server)

    if err := server.Serve(listen); err != nil {
        log.Fatalln(err)
    }
}