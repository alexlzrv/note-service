package note_v1

import (
	"context"

	"github.com/alexlzrv/note-service/internal/mapper"
	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := n.note.CreateNote(ctx, mapper.DescToNote(req.GetNote()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
