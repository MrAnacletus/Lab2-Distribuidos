package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main(){
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	serv := grpc.NewServer()
	pb.RegisterHelloServiceServer(serv, &server{})
	if err = serv.Serve(listener); err != nil{
		panic(err)
	}
}