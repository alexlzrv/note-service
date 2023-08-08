package entity

import (
	"database/sql"
	"time"
)

type Note struct {
	Id        int64        `db:"id"`
	Title     string       `db:"title"`
	Text      string       `db:"text"`
	Author    string       `db:"author"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UpdateNoteInfo struct {
	Id     int64          `db:"id"`
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
}
