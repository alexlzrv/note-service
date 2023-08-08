package mapper

import (
	"database/sql"

	"github.com/alexlzrv/note-service/internal/entity"
	desc "github.com/alexlzrv/note-service/pkg/note_v1"
)

func DescToNote(desc *desc.Note) *entity.Note {
	return &entity.Note{
		Title:  desc.GetTitle(),
		Text:   desc.GetText(),
		Author: desc.GetAuthor(),
	}
}

func NoteToDesc(note *entity.Note) *desc.Note {
	return &desc.Note{
		Title:  note.Title,
		Text:   note.Text,
		Author: note.Author,
	}
}

func DescToNoteInfo(desc *desc.UpdateRequest) *entity.UpdateNoteInfo {
	var title, text, author sql.NullString

	if desc.GetNoteInfo().GetTitle() != nil {
		title.String = desc.GetNoteInfo().GetTitle().Value
		title.Valid = true
	}

	if desc.GetNoteInfo().GetText() != nil {
		text.String = desc.GetNoteInfo().GetText().Value
		text.Valid = true
	}

	if desc.GetNoteInfo().GetAuthor() != nil {
		author.String = desc.GetNoteInfo().GetAuthor().Value
		author.Valid = true
	}
	return &entity.UpdateNoteInfo{
		Id:     desc.GetId(),
		Title:  title,
		Text:   text,
		Author: author,
	}
}
