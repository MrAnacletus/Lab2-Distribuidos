package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MrAnacletus/Lab2-Distribuidos/prueba/proto"
	"google.golang.org/grpc"
)


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

func main(){
	fmt.Println("Iniciando el cliente")
	fmt.Println("Bienvenido al Juego del Calamar")
	fmt.Println("Para comenzar a jugar, presiona enter")
	fmt.Scanln()
	EnviarPeticionJugar()
}