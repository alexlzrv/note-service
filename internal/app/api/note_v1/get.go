package note_v1

import (
	"context"

	"github.com/alexlzrv/note-service/internal/mapper"
	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.note.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: mapper.NoteToDesc(note),
	}, nil
}
