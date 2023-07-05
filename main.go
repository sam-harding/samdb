package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/sam-harding/samdb/protos"
	"google.golang.org/grpc"
)

const (
	SERVER_TYPE  = "tcp"
	SERVER_HOST  = "localhost"
	SERVER_PORT  = "2012"
	MAX_MSG_SIZE = 4096
)

type server struct {
	pb.UnimplementedServerServer
}

func err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
func server_addr(server_host string, server_port string) string {
	return server_host + ":" + server_port
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Println("Get Request: " + in.GetKey())
	return &pb.GetResponse{Value: []byte("This is the value")}, nil
}

func main() {

	server_address := server_addr(SERVER_HOST, SERVER_PORT)

	listener, err := net.Listen(SERVER_TYPE, server_address)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	grpc_server := grpc.NewServer()
	pb.RegisterServerServer(grpc_server, &server{})
	fmt.Println("Listening on " + server_address)

	if err := grpc_server.Serve(listener); err != nil {
		err_panic(err)
	}
}
