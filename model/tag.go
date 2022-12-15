package model

import (
	"database/sql"
	"event-management/db"
)

type TagsMap map[int]Tag

func (tags *TagsMap) FindAll() (err error) {
	var rows *sql.Rows

	sqlStatement := `SELECT * FROM tags WHERE deleted_at IS NULL ORDER BY id`

	rows, err = db.Connection.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var tag = Tag{}
		err = rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

		if err != nil {
			return
		}

		(*tags)[tag.ID] = tag
	}

	return
}

type Tags []Tag

func (tags *Tags) FindAll() (err error) {
	var rows *sql.Rows

	sqlStatement := `SELECT * FROM tags WHERE deleted_at IS NULL ORDER BY id`

	rows, err = db.Connection.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var tag = Tag{}
		err = rows.Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

		if err != nil {
			return
		}

		*tags = append(*tags, tag)
	}

	return
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Timestamps
}

func (tag *Tag) Create() (err error) {
	sqlStatement := `
		INSERT INTO tags (name) VALUES ($1)
		Returning *
	`

	err = db.Connection.
		QueryRow(sqlStatement, tag.Name).
		Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (tag *Tag) Find() (err error) {
	sqlStatement := `SELECT * FROM tags WHERE deleted_at IS NULL AND id = $1`

	err = db.Connection.QueryRow(sqlStatement, tag.ID).
		Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (tag *Tag) FindByName() (err error) {
	sqlStatement := `SELECT * FROM tags WHERE deleted_at IS NULL AND name = $1`

	err = db.Connection.QueryRow(sqlStatement, tag.Name).
		Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (tag *Tag) Update() (err error) {
	sqlStatement := `
		UPDATE tags
		SET name = $2, updated_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
		Returning *
	`

	err = db.Connection.QueryRow(sqlStatement, tag.ID, tag.Name).
		Scan(&tag.ID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (tag *Tag) Delete() error {
	sqlStatement := `
		UPDATE tags
		SET updated_at = NOW(), deleted_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
	`

	_, err := db.Connection.Exec(sqlStatement, tag.ID)
	if err != nil {
		return err
	}

	return nil
}
