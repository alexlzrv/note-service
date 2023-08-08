package note_v1

import (
	"github.com/alexlzrv/note-service/internal/usecase"
	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server
	note *usecase.NoteUseCase
}

func NewNote(note *usecase.NoteUseCase) *Note {
	return &Note{
		note: note,
	}
}
