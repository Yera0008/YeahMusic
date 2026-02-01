package models

type Album struct {
	ID          int64  `json:"id"`
	ArtistID    int64  `json:"artist_id"`
	Title       string `json:"title"`
	CoverURL    string `json:"cover_url"`
	ReleaseYear int    `json:"release_year"`
}
