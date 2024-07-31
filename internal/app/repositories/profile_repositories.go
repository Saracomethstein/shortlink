package repositories

import (
	"database/sql"
	"shortlink/internal/app/models"
)

type IProfileRepository interface {
}

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (p *ProfileRepository) GetUserHistory(login string) ([]models.Link, error) {
	rows, err := p.db.Query("SELECT short_url, original_url FROM urls WHERE login=$1 ORDER BY id DESC", login)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var urls []models.Link
	for rows.Next() {
		var url models.Link
		err := rows.Scan(&url.ShortLink, &url.OriginalLink)
		if err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (p *ProfileRepository) GetLoginFromLog(session_id string) (string, error) {
	var login string
	query := `SELECT login FROM users_log WHERE session_id = $1`
	err := p.db.QueryRow(query, session_id).Scan(&login)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return login, nil
}
