package note_v1

import (
	"context"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := n.note.DeleteNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
