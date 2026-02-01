package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"YeahMusic/internal/services"
)

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (a *App) Register(w http.ResponseWriter, r *http.Request) {
	var req registerReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, 400, "bad json")
		return
	}
	req.Email = strings.TrimSpace(req.Email)
	req.Name = strings.TrimSpace(req.Name)
	if req.Email == "" || req.Password == "" || req.Name == "" {
		writeErr(w, 400, "email/password/name required")
		return
	}

	u, err := a.AuthS.Register(req.Email, req.Password, req.Name)
	if err == services.ErrAlreadyExists {
		writeErr(w, 409, "email already used")
		return
	}
	if mapServiceErr(w, err) {
		return
	}
	writeJSON(w, 201, u)
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, 400, "bad json")
		return
	}
	sec, u, err := a.AuthS.Login(req.Email, req.Password)
	if err != nil {
		writeErr(w, 401, "invalid credentials")
		return
	}
	writeJSON(w, 200, map[string]any{
		"token":      sec.Token,
		"expires_at": sec.ExpiresAt,
		"user":       u,
	})
}
