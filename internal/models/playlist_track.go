package models

import "time"

type PlaylistTrack struct {
	PlaylistID int64     `json:"playlist_id"`
	TrackID    int64     `json:"track_id"`
	AddedAt    time.Time `json:"added_at"`
}
