package server

import "time"

type AuthMessage struct {
	Token string `json:"token"`
}

type ExpirationMessage struct {
	ExpiresAt time.Time `json:"lobby_expires_at"`
}
