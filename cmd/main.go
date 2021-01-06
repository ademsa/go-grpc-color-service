package main

import (
	"flag"
	"go-grpc-color-service/pkg/server"
	"go-grpc-color-service/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func main() {
	port := flag.Int("port", 3002, "Port")

	flag.Parse()

	if *port < 3000 {
		log.Println("[Input] Validation Error", "Please specify port >= 3000")
		return
	}

	log.Println("Go gRPC Color Service")
	log.Println("Port:", *port)

	l, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	service.RegisterColorServiceServer(s, &server.ColorServer{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
