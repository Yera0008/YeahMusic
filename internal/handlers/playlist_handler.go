package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type createPlaylistReq struct {
	Title string `json:"title"`
}

func (a *App) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	u := userFromCtx(r)
	var req createPlaylistReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, 400, "bad json")
		return
	}
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		writeErr(w, 400, "title required")
		return
	}
	p := a.PlayS.Create(u.ID, req.Title)
	writeJSON(w, 201, p)
}

func (a *App) ListPlaylists(w http.ResponseWriter, r *http.Request) {
	u := userFromCtx(r)
	writeJSON(w, 200, a.PlayS.List(u.ID))
}

type addTrackReq struct {
	TrackID int64 `json:"track_id"`
}

func (a *App) PlaylistSubroutes(w http.ResponseWriter, r *http.Request) {
	u := userFromCtx(r)

	path := strings.TrimPrefix(r.URL.Path, "/api/playlists/")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		writeErr(w, 404, "not found")
		return
	}

	playlistID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		writeErr(w, 400, "playlist id must be int")
		return
	}

	if r.Method == http.MethodPost && len(parts) == 2 && parts[1] == "tracks" {
		var req addTrackReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeErr(w, 400, "bad json")
			return
		}
		if req.TrackID <= 0 {
			writeErr(w, 400, "track_id required")
			return
		}
		err := a.PlayS.AddTrack(u.ID, playlistID, req.TrackID)
		if mapServiceErr(w, err) {
			return
		}
		writeJSON(w, 200, map[string]any{"status": "added"})
		return
	}

	if r.Method == http.MethodDelete && len(parts) == 3 && parts[1] == "tracks" {
		trackID, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			writeErr(w, 400, "track id must be int")
			return
		}
		err = a.PlayS.RemoveTrack(u.ID, playlistID, trackID)
		if mapServiceErr(w, err) {
			return
		}
		writeJSON(w, 200, map[string]any{"status": "removed"})
		return
	}

	writeErr(w, 404, "not found")
}
