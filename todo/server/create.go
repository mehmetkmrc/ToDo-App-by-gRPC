package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(s *Server) CreateEvent(ctx context.Context, in *pb.Event) (*pb.EventId, error){
	log.Printf("CreateEvent uyandırıldı : %v\n", in)

	// Execute SQL query to insert the event into the database
	query := "INSERT INTO events (author_id, title, content) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	rows, err :=db.DB.Query(query, in.AuthorId, in.Title, in.Content)
	//err := db.DB.QueryRowContext(ctx, query, in.AuthorId, in.Title, in.Content).Scan(&id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v\n", err),
		)
	}
	defer rows.Close()

	// Convert the inserted ID to hexadecimal string
	idHex := fmt.Sprintf("%x", id)

	return &pb.EventId{
		Id: idHex,
	}, nil

}