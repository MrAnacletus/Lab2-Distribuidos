package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Peticion recibida, aceptando juego")
	return &pb.HelloReply{Message: "Juego aceptado"}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ")
	return &pb.HelloReply{Message: "Hello "}, nil
}

func (s *server) GetJugadas(ctx context.Context, in *pb.Jugadas) (*pb.HelloReply, error) {
	fmt.Println("Jugadas recibidas ")
	return &pb.HelloReply{Message: "Jugadas recibidas, gracias"}, nil
}

func EnviarJugadas(){
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al server: %v",err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	fmt.Println("Enviando jugadas")
	resp, err := client.GetJugadas(context.Background(), &pb.Jugadas{})
	if err != nil {
		log.Fatalf("No se pudo enviar el mensaje: %v",err)
	}
	fmt.Println("Respuesta: ",resp.Message)
}

func ServidorCliente(){
	listener , err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}
}

func main(){
	ServidorCliente()
}