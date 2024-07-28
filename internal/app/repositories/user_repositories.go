package repositories

import "database/sql"

type UserRepository interface {
	CheckUserExistsByLogin(login string) (bool, error)
	FindUserByLogin(login string) (*User, error)
	CreateUser(use *User) error
}

type User struct {
	ID       int
	Login    string
	Password string
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CheckUserExistsByLogin(login string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)`
	err := r.db.QueryRow(query, login).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *userRepository) FindUserByLogin(login string) (*User, error) {
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

func (r *userRepository) CreateUser(user *User) error {
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
