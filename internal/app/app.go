package app

import (
	"github.com/alexlzrv/note-service/internal/app/api/note_v1"
	"google.golang.org/grpc"
	"log"
	"net"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

const port = ":50051"

func InitServer() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
