package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
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

func (s *server) GetJugada(ctx context.Context, in *pb.Jugada) (*pb.Resultado, error) {
	// Enviarla a NameNode
	fmt.Println("Jugadas recibidas, jugada:" + fmt.Sprint(in.GetJugada()) + " del jugador: " + fmt.Sprint(in.GetID()))
	conn, err := grpc.Dial("10.6.40.219:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server NameNode: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewHelloServiceClient(conn)
	//Se envia la jugada apra que sea escrita en el archivo
	_, err = serviceClient.GetJugada(context.Background(), &pb.Jugada{Jugada: in.GetJugada(), ID: in.GetID()})
	if err != nil {
		log.Fatalf("No se pudo enviar la jugada: %v",err)
	}
	//Lider debe elegir un numero entre 6 y 10
	JugadaLider := rand.Intn(4) + 6
	if JugadaLider <= int(in.GetJugada()){
		//Notificar a pozo que alguien murio por rabbitmq
		return &pb.Resultado{ID: in.GetID(),Estado: 0}, nil
	}
	return &pb.Resultado{ID: in.GetID(),Estado: 1}, nil
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
