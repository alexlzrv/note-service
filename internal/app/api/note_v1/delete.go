package note_v1

import (
	"context"
	"fmt"

	desc "github.com/alexlzrv/note-service/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("Delete")

	return &emptypb.Empty{}, nil
}
