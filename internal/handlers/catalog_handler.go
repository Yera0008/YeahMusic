package handlers

import (
	"net/http"
	"strconv"
)

func (a *App) ListArtists(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, a.CatalogS.ListArtists())
}

func (a *App) ListAlbums(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("artist_id")
	var artistID int64
	if q != "" {
		n, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			writeErr(w, 400, "artist_id must be int")
			return
		}
		artistID = n
	}
	writeJSON(w, 200, a.CatalogS.ListAlbums(artistID))
}

func (a *App) ListTracks(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("album_id")
	var albumID int64
	if q != "" {
		n, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			writeErr(w, 400, "album_id must be int")
			return
		}
		albumID = n
	}
	writeJSON(w, 200, a.CatalogS.ListTracks(albumID))
}
