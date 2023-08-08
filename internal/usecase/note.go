package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/alexlzrv/note-service/internal/entity"
)

type NoteUseCase struct {
	repo    NoteRepository
	timeout time.Duration
}

func NewNoteUseCase(repo NoteRepository, timeout time.Duration) *NoteUseCase {
	return &NoteUseCase{
		repo:    repo,
		timeout: timeout,
	}
}

func (n *NoteUseCase) CreateNote(ctx context.Context, note *entity.Note) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, n.timeout)
	defer cancel()

	id, err := n.repo.CreateNote(ctx, note)
	if err != nil {
		return 0, fmt.Errorf("cannot create note: %w", err)
	}

	return id, nil
}

func (n *NoteUseCase) GetNote(ctx context.Context, id int64) (*entity.Note, error) {
	ctx, cancel := context.WithTimeout(ctx, n.timeout)
	defer cancel()

	return n.repo.GetNote(ctx, id)
}

func (n *NoteUseCase) UpdateNote(ctx context.Context, updateNoteInfo *entity.UpdateNoteInfo) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, n.timeout)
	defer cancel()

	return n.repo.UpdateNote(ctx, updateNoteInfo)
}

func (n *NoteUseCase) DeleteNote(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, n.timeout)
	defer cancel()

	return n.repo.DeleteNote(ctx, id)
}
