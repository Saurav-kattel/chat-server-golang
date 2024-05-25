package models

type CreateRoomModel struct {
	IsGroup bool     `json:"isGroup"`
	Users   []string `json:"users,omitempty"`
}
