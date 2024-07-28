package repositories

import (
	"database/sql"
)

type LinkRepository interface {
	GetOriginalLink(shortLink string) (*Link, error)
	GetShortLink(originalLink string) (*Link, error)
	CreateShortLink(link *Link) error
	CheckLinkExistByOriginal(originalLink string) (bool, error)
}

type Link struct {
	ID          int64
	ShortCode   string
	OriginalURL string
}

type linkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *linkRepository {
	return &linkRepository{db: db}
}

func (r *linkRepository) GetOriginalLink(shortLink string) (*Link, error) {
	link := &Link{}
	query := "SELECT original_url FROM urls WHERE short_url = $1"
	err := r.db.QueryRow(query, shortLink).Scan(&link.ID, &link.ShortCode, &link.OriginalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return link, nil
}

func (r *linkRepository) GetShortLink(originalLink string) (*Link, error) {
	link := &Link{}
	query := "SELECT short_url FROM urls WHERE original_url = $1"
	err := r.db.QueryRow(query, originalLink).Scan(&link.ID, &link.ShortCode, &link.OriginalURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return link, nil
}

func (r *linkRepository) CreateShortLink(link *Link) error {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE short_url=$1 AND original_url=$2)`
	err := r.db.QueryRow(query, link.ShortCode, link.OriginalURL).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		query = `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`
		_, err := r.db.Exec(query, link.ShortCode, link.OriginalURL)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *linkRepository) CheckLinkExistByOriginal(originalLink string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM urls WHERE original_url=$1)`
	err := r.db.QueryRow(query, originalLink).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
