package main

import (
	"log"
	"net"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"


type Server struct{
	pb.EventServiceServer
}
func main(){
	var err error
	db.InitDB()
	

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
	
}

