package main

import (
	"context"
	"math/rand"
	"math"
	"fmt"
	"log"
	"net"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)
type server struct{
	pb.UnimplementedLiderServiceServer
}
type Jugada struct {
	ID int32
	Jugada int32
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Peticion recibida, aceptando juego")

	return &pb.HelloReply{Message: "Juego aceptado"}, nil
}

func (s *server) SendJugada(ctx context.Context, in *pb.Jugada) (*pb.Resultado, error) {
	// Enviarla a NameNode
	fmt.Println("Jugadas recibidas, jugada:" + fmt.Sprint(in.GetJugada()) + " del jugador: " + fmt.Sprint(in.GetID()))
	conn, err := grpc.Dial("10.6.40.219:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server NameNode: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewNameNodeServiceClient(conn)
	//Se envia la jugada para que sea escrita en el archivo
	_, err = serviceClient.SendJugada(context.Background(), &pb.Jugada{Jugada: in.GetJugada(), ID: in.GetID()})
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

func (s *server) SendJugada2(ctx context.Context, in *pb.Jugada) (*pb.Resultado, error) {
	//Recibir las jugadas
	T1 := in.GetID()
	T2 := in.GetJugada()

	JugadaLider := (rand.Intn(4) + 1) % 2

	if JugadaLider == 0{
		if  T1%2 == 0 && T2%2 == 0{
			return &pb.Resultado{ID: 1,Estado: 1}, nil
		}else if T1%2 == 1 && T2%2 == 1{
			Decididor := rand.Intn(1)
			if Decididor == 0{
				return &pb.Resultado{ID: 1,Estado: 0}, nil
			}else{
				return &pb.Resultado{ID: 0,Estado: 1}, nil
			}
		}else if T1%2 == 0 && T2%2 == 1{
			return &pb.Resultado{ID: 1,Estado: 0}, nil
		}else{
			return &pb.Resultado{ID: 0,Estado: 1}, nil
		}
	}else{
		if  T1%2 == 1 && T2%2 == 1{
			return &pb.Resultado{ID: 1,Estado: 1}, nil
		}else if T1%2 == 0 && T2%2 == 0{
			Decididor := rand.Intn(1)
			if Decididor == 0{
				return &pb.Resultado{ID: 1,Estado: 0}, nil
			}else{
				return &pb.Resultado{ID: 0,Estado: 1}, nil
			}
		}else if T1%2 == 0 && T2%2 == 1{
			return &pb.Resultado{ID: 0,Estado: 1}, nil
		}else{
			return &pb.Resultado{ID: 1,Estado: 0}, nil
		}
	}
}

func (s *server) SendJugada3 (ctx context.Context, in *pb.Jugada3) (*pb.Resultado, error) {
	//Recibir las jugadas
	ID1 := in.ID1
	ID2 := in.ID2
	jugada1 := in.Jugada1
	jugada2 := in.Jugada2

	//Lider debe elegir un numero entre 1 y 10
	JugadaLider := rand.Intn(10) + 1
	//Calcular distancias con el lider
	Distancia1 := math.Abs(float64(JugadaLider - int(jugada1)))
	Distancia2 := math.Abs(float64(JugadaLider - int(jugada2)))
	//Decidir al ganador
	if Distancia1 < Distancia2{
		return &pb.Resultado{ID: ID1,Estado: 1}, nil
	}
	return &pb.Resultado{ID: ID2,Estado: 1}, nil

}

func abs(i int) {
	panic("unimplemented")
}


func (s *server) RequestPozo(ctx context.Context, in *pb.RequestPozoActual) (*pb.ResponsePozoActual, error) {
	// Enviarla a Pozo
	fmt.Println("Peticion recibida, enviando peticion al servidor Pozo")
	conn, err := grpc.Dial("10.6.40.220:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el server Pozo: %v",err)
	}
	defer conn.Close()

	//Se crea un cliente para la conexion
	serviceClient := pb.NewPozoServiceClient(conn)
	//Se envia la peticion al servidor pozo
	res, err := serviceClient.RequestPozo(context.Background(), &pb.RequestPozoActual{Pozo: 1})
	if err != nil {
		log.Fatalf("No se pudo enviar la peticion: %v",err)
	}
	return res, nil
}

func ServidorCliente(){
	listener , err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterLiderServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("No se pudo iniciar el server: %v",err)
	}
}

func main(){
	ServidorCliente()
	
}