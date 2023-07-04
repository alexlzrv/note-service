package note_v1

import (
	"context"
	"fmt"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update")
	fmt.Println("title:", req.Note.GetTitle())
	fmt.Println("text:", req.Note.GetText())
	fmt.Println("author:", req.Note.GetAuthor())

	return &emptypb.Empty{}, nil
}
