package handlers

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/chat-server/internal/models"
)

func createRoom(db *sqlx.DB, data *models.CreateRoomModel) error {
	_, err := db.Exec("INSERT INTO rooms(is_group) VALUES($1)", data.IsGroup)
	if err != nil {
		return err
	}
	return nil
}

func createRoomUsers(db *sqlx.DB, data *models.RoomUsersPayload) error {
	_, err := db.Exec("INSERT INTO room_users(user_id) VALUES($0)", data.UserId)
	if err != nil {
		return err
	}
	return nil
}

func getRoomUsers(db *sqlx.DB, roomId string) (*[]models.RoomUsers, error) {
	var users []models.RoomUsers

	err := db.Select(&users, "SELECT * FROM room_users WHERE room_id = $1", roomId)
	if err != nil {
		return nil, err
	}
	return &users, nil
}
