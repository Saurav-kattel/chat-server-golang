package handlers

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/ws-server/internal/models"
)

func CreateRoom(db *sqlx.DB, data *models.CreateRoomModel) error {
	_, err := db.Exec("INSERT INTO rooms(is_group) VALUES($1)", data.IsGroup)
	if err != nil {
		return err
	}
	return nil
}
