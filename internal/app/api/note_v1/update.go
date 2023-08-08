package note_v1

import (
	"context"

	"github.com/alexlzrv/note-service/internal/mapper"
	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := n.note.UpdateNote(ctx, mapper.DescToNoteInfo(req))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
