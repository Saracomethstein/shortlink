package models

type LoginRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type CreateShortLinkRequest struct {
	OriginalURL string `json:"url"`
}

type CreateShortLinkResponse struct {
	ShortCode string `json:"shortenedUrl"`
}
