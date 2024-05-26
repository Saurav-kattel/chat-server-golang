package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

func Route(db *sqlx.DB) *http.ServeMux {
	rotue := http.ServeMux{}
	rotue.HandleFunc("/api/v1/room/init", CreateRoomHandler(db))
	return &rotue
}
