package main

import (
	"context"
	"log"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func(s *Server) DeleteEvent(ctx context.Context, in *pb.EventId) (*emptypb.Empty, error){
	log.Printf("Deleting çalıştırıldı: %v\n", in)


	query := "DELETE FROM events WHERE id =$1 "
	stmt, err :=db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(in.Id)
	return &emptypb.Empty{}, nil
}