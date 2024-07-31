package repositories

import (
	"database/sql"
)

type IUserRepository interface {
	CheckUserExistsByLogin(login string) (bool, error)
	CheckUserExists(user User) (bool, error)
	FindUserByLogin(login string) (*User, error)
	CreateUser(user *User) error
}

type User struct {
	ID       int
	Login    string
	Password string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckUserExistsByLogin(login string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)`
	err := r.db.QueryRow(query, login).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *UserRepository) FindUserByLogin(login string) (*User, error) {
	user := &User{}
	query := `SELECT id, login, password FROM users WHERE login = $1`
	err := r.db.QueryRow(query, login).Scan(&user.ID, &user.Login, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user *User) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login=$1 AND password=$2)`
	err := r.db.QueryRow(query, user.Login, user.Password).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		query = `INSERT INTO users ( login, password) VALUES ($1, $2)`
		_, err := r.db.Exec(query, user.Login, user.Password)

		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) CheckUserExists(user User) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login=$1 AND password=$2)`
	err := r.db.QueryRow(query, user.Login, user.Password).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *UserRepository) CreateUserLogging(login, session_id string) error {
	query := `INSERT INTO users_log ( login, session_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, login, session_id)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindUserLog(login, session_id string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users_log WHERE login=$1 AND session_id=$2)`
	var err = r.db.QueryRow(query, login, session_id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
