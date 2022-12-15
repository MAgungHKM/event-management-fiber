package model

import (
	"database/sql"
	"event-management/db"
)

type UsersMap map[int]User

func (usersMap *UsersMap) FindAll() (err error) {
	var rows *sql.Rows

	sqlStatement := `SELECT * FROM users WHERE deleted_at IS NULL ORDER BY id`

	rows, err = db.Connection.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		if err != nil {
			return
		}

		(*usersMap)[user.ID] = user
	}

	return
}

type Users []User

func (users *Users) FindAll() (err error) {
	var rows *sql.Rows

	sqlStatement := `SELECT * FROM users WHERE deleted_at IS NULL ORDER BY id`

	rows, err = db.Connection.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		if err != nil {
			return
		}

		*users = append(*users, user)
	}

	return
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Secret   string `json:"-"`
	Timestamps
}

func (user *User) Create() (err error) {
	sqlStatement := `
		INSERT INTO users (name, username, email, secret) VALUES ($1, $2, $3, $4)
		Returning *
	`

	err = db.Connection.
		QueryRow(sqlStatement, user.Name, user.Username, user.Email, user.Secret).
		Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (user *User) Find() (err error) {
	sqlStatement := `SELECT * FROM users WHERE deleted_at IS NULL AND id = $1`

	err = db.Connection.QueryRow(sqlStatement, user.ID).
		Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (user *User) FindByUsername() (err error) {
	sqlStatement := `SELECT * FROM users WHERE deleted_at IS NULL AND username = $1`

	err = db.Connection.QueryRow(sqlStatement, user.Username).
		Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (user *User) Update() (err error) {
	sqlStatement := `
		UPDATE users
		SET name = $2, username = $3, email = $4, secret = $5, updated_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
		Returning *
	`

	err = db.Connection.QueryRow(sqlStatement, user.ID, user.Name, user.Username, user.Email, user.Secret).
		Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Secret, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (user *User) Delete() error {
	sqlStatement := `
		UPDATE users
		SET updated_at = NOW(), deleted_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
	`

	_, err := db.Connection.Exec(sqlStatement, user.ID)
	if err != nil {
		return err
	}

	return nil
}
