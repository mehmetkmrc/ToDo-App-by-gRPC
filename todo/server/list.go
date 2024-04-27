package main

import (
	"log"

	"github.com/mehmetkmrc/ToDoApp/todo/db"
	pb "github.com/mehmetkmrc/ToDoApp/todo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)


func (s *Server) ListEvents(in *emptypb.Empty, stream pb.EventService_ListEventsServer) error {
	log.Println("ListEvents called")

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var event pb.Event
		err := rows.Scan(&event.Id, &event.AuthorId, &event.Title, &event.Content)
		if err != nil {
			return err
		}
		if err := stream.Send(&event); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
