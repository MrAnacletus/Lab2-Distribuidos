package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)

type Jugada struct {
	ID int32
	jugada int32
}

func EnviarPeticionJugar(){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewHelloServiceClient(conn)
	//Se envia la petición de jugar al servidor
	res, err := serviceCLient.SayHello(context.Background(), &pb.HelloRequest{Name: "Preparandoce para iniciar el juego!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Message)
}

func EnviarJugada(J Jugada){
	//Se establece la conexión con el servidor
	conn, err := grpc.Dial("10.6.40.218:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse con el servidor Lider: %v", err)
	}
	defer conn.Close()

	//Se crea un cliente para la comunicación con el servidor
	serviceCLient := pb.NewHelloServiceClient(conn)
	//Se envia la Jugada al servidor
	res, err := serviceCLient.GetJugada(context.Background(), &pb.Jugada{ID: J.ID, Jugada: J.jugada})
	if err != nil {
		log.Fatalf("Error al enviar la Jugada: %v", err)
	}
	fmt.Println(res.Message)
}

func juego1(){
	for i := 0; i < 4; i++ {
		fmt.Println("Elija un numero entre 1 y 10, recuerde que sus numeros deben sumar 21.")
		var numero1 int
		fmt.Scanln(&numero1)
		j = Jugada{ID: 1, jugada: int32(numero1)}
	}
}

func main(){
	fmt.Println("Iniciando el cliente")
	fmt.Println("Bienvenido al Juego del Calamar")
	fmt.Println("Para comenzar a jugar, presiona enter")
	fmt.Scanln()
	EnviarPeticionJugar()
	fmt.Println("El juego ha comenzado")
	fmt.Println("Selecciona que vas a hacer:")
	fmt.Println("1. Jugar")
	fmt.Println("2. Ver pozo")
	var opcion int
	fmt.Scanln(&opcion)
	if opcion == 1 {
		juego1()
	}
}