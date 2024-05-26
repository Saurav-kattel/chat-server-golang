package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"x-clone.com/chat-server/internal/models"
)

// Mock utility functions
func MockJsonDecoder[T any](r *http.Request) (*T, error) {
	var t T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	return &t, err
}

func MockResponseJson(w http.ResponseWriter, status int, payload interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func TestCreateRoomHandler(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create room and room users data
	roomData := &models.CreateRoomModel{
		IsGroup: true,
		Users:   []string{"user1", "user2"},
	}

	// Mock expected database actions
	mock.ExpectExec("INSERT INTO rooms").WithArgs(roomData.IsGroup).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO room_users").WithArgs("user1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO room_users").WithArgs("user2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a request to pass to our handler
	reqBody, _ := json.Marshal(roomData)
	req := httptest.NewRequest("POST", "/create-room", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create handler and serve request
	handler := CreateRoomHandler(sqlxDB)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Check the response body is what we expect
	expected := `{"status":200,"res":"room created successfully"}`
	assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
