package usecase

import (
	"context"

	"github.com/alexlzrv/note-service/internal/entity"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, note *entity.Note) (int64, error)
	GetNote(ctx context.Context, id int64) (*entity.Note, error)
	UpdateNote(ctx context.Context, updateNoteInfo *entity.UpdateNoteInfo) (int64, error)
	DeleteNote(ctx context.Context, id int64) error
}
