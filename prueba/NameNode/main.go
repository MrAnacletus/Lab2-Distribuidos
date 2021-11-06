package main

import(
	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
	"context"
	"fmt"
	"net"
)

type server struct{
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Enviado: ")
	return &pb.HelloReply{Message: "Hello "}, nil
}

func (s *server) GetJugada(ctx context.Context, in *pb.Jugada) (*pb.HelloReply, error) {
	fmt.Println("Jugada recibida, Jugada: " + fmt.Sprint(in.Jugada) + " Jugador: " + fmt.Sprint(in.ID))
	return &pb.HelloReply{Message: "Jugadas recibidas, gracias"}, nil
}

func ServidorNameNode(){
	listener , err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("Error al iniciar el servidor")
		return
	}
	fmt.Println("Servidor iniciado")
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err:= s.Serve(listener); err != nil{
		fmt.Println("Error al iniciar el servidor")
		return
	}
}


func main(){
	ServidorNameNode()
}
