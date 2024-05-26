package models

type CreateRoomModel struct {
	IsGroup bool     `json:"isGroup"`
	Users   []string `json:"users,omitempty"`
}

type RoomUsersPayload struct {
	UserId string `json:"userId"`
}

type RoomUsers struct {
	Id     string `json:"id" db:"id"`
	UserId string `json:"userId" db:"user_id"`
}

type Message struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status int     `json:"status"`
	Res    Message `json:"res"`
}

type SuccessResponse struct {
	Status int `json:"status"`
	Res    any `json:"res"`
}
