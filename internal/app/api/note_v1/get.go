package note_v1

import (
	"context"
	"fmt"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get")
	fmt.Println("Id:", req.GetId())

	note := desc.Note{
		Title:  "title2",
		Text:   "text2",
		Author: "author2",
	}

	return &desc.GetResponse{
		Note: &note,
	}, nil
}
