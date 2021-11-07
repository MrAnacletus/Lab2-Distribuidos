package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedNameNodeServiceServer
}

func (s *server) SendJugada(ctx context.Context, in *pb.Jugada) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Jugada Recibida"}, nil
}

func (s *server) GetJugada(ctx context.Context, in *pb.RequestJugadas) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Jugada Recibida"}, nil
}
func ServidorNameNode(){
	listener , err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("Error al iniciar el servidor")
		return
	}
	fmt.Println("Servidor NameNode Iniciado")
	s := grpc.NewServer()
	pb.RegisterNameNodeServiceServer(s, &server{})
	if err:= s.Serve(listener); err != nil{
		fmt.Println("Error al iniciar el servidor")
		return
	}
}


func main(){
	ServidorNameNode()
}
