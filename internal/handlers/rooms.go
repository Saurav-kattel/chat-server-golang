package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/chat-server/internal/models"
	"x-clone.com/chat-server/utils"
)

// saves a roomid in db

func createRoom(db *sqlx.DB, data *models.CreateRoomModel) (string, error) {
	var id string
	err := db.QueryRowx("INSERT INTO rooms(is_group) VALUES($1) RETURNING id", data.IsGroup).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// creates a users for room
func createRoomUsers(db *sqlx.Tx, userId, roomId string) error {
	_, err := db.Exec("INSERT INTO room_users(user_id,room_id) VALUES($1,$2)", userId, roomId)
	return err
}

// retrives all the users for a room
func GetRoomUsers(db *sqlx.DB, roomId string) (*[]models.RoomUsers, error) {
	var users []models.RoomUsers

	err := db.Select(&users, "SELECT * FROM room_users WHERE room_id = $1", roomId)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// a function to handle  room creation
func CreateRoomHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			utils.ResponseJson(w, http.StatusMethodNotAllowed,
				models.ErrorResponse{
					Status: http.StatusMethodNotAllowed,
					Res: models.Message{
						Message: "invalid method",
					},
				})
		}

		// custom jsonDecoder to decode json
		data, err := utils.JsonDecoder[models.CreateRoomModel](r)
		if err != nil {
			utils.ResponseJson(w, http.StatusInternalServerError,
				models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
			return
		}

		roomId, err := createRoom(db, data)
		if err != nil {
			utils.ResponseJson(w, http.StatusInternalServerError,
				models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
			return
		}

		tx, err := db.Beginx()
		if err != nil {
			utils.ResponseJson(w, http.StatusInternalServerError,
				models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
			return

		}
		for _, users := range data.Users {
			if err := createRoomUsers(tx, users, roomId); err != nil {
				tx.Rollback()
				utils.ResponseJson(w, http.StatusInternalServerError,
					models.ErrorResponse{
						Status: http.StatusInternalServerError,
						Res: models.Message{
							Message: err.Error(),
						},
					})
				return

			}
		}
		if commitErr := tx.Commit(); commitErr != nil {

			utils.ResponseJson(w, http.StatusInternalServerError,
				models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: commitErr.Error(),
					},
				})
			return

		}

		utils.ResponseJson(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "room created successfully",
		})
	}
}
