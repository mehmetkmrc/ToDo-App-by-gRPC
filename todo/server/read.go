package main

import (
	"context"
	"log"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
)

func(s *Server) ReadEvent(ctx context.Context, in *pb.EventId) (*pb.Event, error){
	log.Printf("Reading was invoked with: %v\n", in)

	query :="SELECT *FROM events WHERE id = $1 "
	row := db.DB.QueryRow(query, in.Id)

	var event pb.Event
	err := row.Scan(&event.Id, &event.AuthorId, &event.Title, &event.Content)

	//Hata mesajı gelecek
	if err != nil {
        log.Printf("Error reading event from database: %v\n", err)
        return nil, err // Hata mesajını döndür
    }

	return &event, nil
	
}