package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexlzrv/note-service/internal/entity"
	"github.com/alexlzrv/note-service/pkg/postgres"
)

const (
	noteTable = "note"
)

type NoteRepo struct {
	db *postgres.Postgres
}

func NewNoteRepo(db *postgres.Postgres) *NoteRepo {
	return &NoteRepo{
		db: db,
	}
}

func (n *NoteRepo) CreateNote(ctx context.Context, note *entity.Note) (int64, error) {
	builder := n.db.Builder.
		Insert(noteTable).
		Columns("title, text, author").
		Values(note.Title, note.Text, note.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("createNote, error with query: %w", err)
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("createNote, error with scan: %w", err)
	}

	return id, nil
}

func (n *NoteRepo) GetNote(ctx context.Context, id int64) (*entity.Note, error) {
	builder := n.db.Builder.
		Select("id, title, text, author, created_at, updated_at").
		From(noteTable).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := n.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var note entity.Note
	rows.Next()
	err = rows.Scan(&note.Id, &note.Title, &note.Text, &note.Author,
		&note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (n *NoteRepo) UpdateNote(ctx context.Context, updateNoteInfo *entity.UpdateNoteInfo) (int64, error) {
	builder := n.db.Builder.
		Update(noteTable).
		Set("updated_at", time.Now())

	if updateNoteInfo.Title.Valid {
		builder = builder.Set("title", updateNoteInfo.Title.String)
	}
	if updateNoteInfo.Text.Valid {
		builder = builder.Set("text", updateNoteInfo.Text.String)
	}
	if updateNoteInfo.Author.Valid {
		builder = builder.Set("author", updateNoteInfo.Author.String)
	}

	builder = builder.Where(sq.Eq{"id": updateNoteInfo.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return updateNoteInfo.Id, nil
}

func (n *NoteRepo) DeleteNote(ctx context.Context, id int64) error {
	builder := n.db.Builder.
		Delete(noteTable).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = n.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
