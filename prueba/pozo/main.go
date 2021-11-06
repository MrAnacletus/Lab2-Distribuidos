package main

import (
	"net"
	"context"
	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
	"log"

)


type server struct{
	pb.UnimplementedPozoServiceServer
}

var POZOACTUAL int32


func (s *server) RequestPozo(ctx context.Context, req *pb.RequestPozoActual) (*pb.ResponsePozoActual, error) {
	return &pb.ResponsePozoActual{
		Pozo: POZOACTUAL,
	}, nil
}

func ServidorPozo(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPozoServiceServer(s, &server{})
	if err:= s.Serve(listener); err != nil {
		log.Fatalf("No se pudo iniciar el servidor %v", err)
	}
}

func main(){
	ServidorPozo()
}