package repository

import (
	"dogma_test_task/internal"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserPostgres struct {
	db sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: *db}
}

func (r *UserPostgres) GetUserList() ([]internal.User, error) {
	var users []internal.User

	query := "SELECT * FROM users"
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) AddUser(user internal.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createQuery := "INSERT INTO users (username, email) values ($1, $2) RETURNING id"

	row := tx.QueryRow(createQuery, user.Username, user.Email)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *UserPostgres) GetUserById(id int) (internal.User, error) {
	var user internal.User
	query := "SELECT * FROM users WHERE id=$1"
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UserPostgres) UpdateUser(id int, input internal.User) (internal.User, error) {
	var user internal.User

	keys := make([]string, 0)
	values := make([]interface{}, 0)

	values = append(values, id)

	if input.Email != "" {
		keys = append(keys, "email")
		values = append(values, input.Email)
	}

	if input.Username != "" {
		keys = append(keys, "username")
		values = append(values, input.Username)
	}

	sets := make([]string, 0)
	for i, v := range keys {
		sets = append(sets, fmt.Sprintf("%s=$%d", v, i+2))
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$1", strings.Join(sets, ", "))
	_, err := r.db.Exec(query, values...)

	if err != nil {
		return user, err
	}

	return r.GetUserById(id)
}


func (r *UserPostgres) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	return err
}