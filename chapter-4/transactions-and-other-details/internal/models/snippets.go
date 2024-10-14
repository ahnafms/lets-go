package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      string
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB         *sql.DB
	InsertStmt *sql.Stmt
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	// stmt := "SELECT id, title, content, created expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	// row := m.DB.QueryRow(stmt, id)
	//
	// var s Snippet
	//
	// err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Expires)

	//another approach
	//
	var s Snippet
	err := m.DB.QueryRow("SELECT id, content, title, created, expires FROM snippets WHERE id = ?", id).Scan(&s.ID, &s.Content, &s.Title, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, content, title, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err := rows.Scan(&s.ID, &s.Content, &s.Title, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

func (m *SnippetModel) ExampleTransaction() error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO...")

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE ...")

	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func NewExampleModel(db *sql.DB) (*SnippetModel, error) {
	insertStmt, err := db.Prepare("INSERT INTO ...")
	if err != nil {
		return nil, err
	}

	return &SnippetModel{DB: db, InsertStmt: insertStmt}, nil
}

func (m *SnippetModel) InsertPrepare(args ...interface{}) error {
	_, err := m.InsertStmt.Exec(args...)

	return err
}
