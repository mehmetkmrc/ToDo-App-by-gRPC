package main

import (
	"context"
	"log"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func(s *Server) UpdateEvent(ctx context.Context, in *pb.Event) (*emptypb.Empty, error){
	log.Printf("updateEvent çalıştırıldı : %v\n", in)
  
	query := `
	UPDATE events
	SET author_id = $1, title = $2, content = $3
	WHERE id = $4
	`
  
	stmt, err := db.DB.Prepare(query)
  
	if err != nil {
	  return nil, err
	}
	defer stmt.Close()
  
	// Use the data from the incoming event
	_, err = stmt.Exec(in.AuthorId, in.Title, in.Content, in.Id)
  
	if err != nil {
	  return nil, err
	}
  
	return &emptypb.Empty{}, nil
  }
  