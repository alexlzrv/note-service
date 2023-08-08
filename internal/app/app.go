package app

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/alexlzrv/note-service/config"
	"github.com/alexlzrv/note-service/internal/app/api/note_v1"
	"github.com/alexlzrv/note-service/internal/usecase"
	"github.com/alexlzrv/note-service/internal/usecase/repository"
	"github.com/alexlzrv/note-service/pkg/postgres"
	"google.golang.org/grpc"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

const port = ":50051"

func RunServer(cfg *config.Config) {

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	repo := repository.NewNoteRepo(pg)

	note := usecase.NewNoteUseCase(repo, 3*time.Second)

	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote(note))

	if err = s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
