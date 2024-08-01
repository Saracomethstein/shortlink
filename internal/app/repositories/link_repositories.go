package repositories

import (
	"database/sql"
)

type ILinkRepository interface {
	GetOriginalLink(shortLink string) (*Link, error)
	GetShortLink(originalLink string) (*Link, error)
	CreateShortLink(link *Link) error
	CheckLinkExistByOriginal(originalLink string) (bool, error)
	GetLoginFromLog(session_id string) (string, error)
}

type Link struct {
	ID          int64
	ShortCode   string
	OriginalURL string
}

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) GetOriginalLink(shortLink string) (string, error) {
	link := &Link{}
	query := "SELECT original_url FROM urls WHERE short_url = $1"
	err := r.db.QueryRow(query, shortLink).Scan(&link.OriginalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return link.OriginalURL, nil
}

func (r *LinkRepository) GetShortLink(originalLink string) (string, error) {
	link := &Link{}
	query := "SELECT short_url FROM urls WHERE original_url = $1"
	err := r.db.QueryRow(query, originalLink).Scan(&link.ShortCode)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return link.ShortCode, nil
}

func (r *LinkRepository) CreateShortLink(login, shortLink, originalLink string) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE short_url=$1 AND original_url=$2)`
	err := r.db.QueryRow(query, shortLink, originalLink).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		query = `INSERT INTO urls (login, short_url, original_url) VALUES ($1, $2, $3)`
		_, err := r.db.Exec(query, login, shortLink, originalLink)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *LinkRepository) CheckLinkExistByOriginal(originalLink string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE original_url=$1)`
	err := r.db.QueryRow(query, originalLink).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r *LinkRepository) GetLoginFromLog(session_id string) (string, error) {
	var login string
	query := `SELECT login FROM users_log WHERE session_id = $1`
	err := r.db.QueryRow(query, session_id).Scan(&login)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return login, nil
}
