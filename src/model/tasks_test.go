package model

import (
	"TravelService/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var taskMockErr error

func TestCreateTask(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, taskMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	var status bool

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	task := Task{
		ID,
		UserID,
		"TaskOne",
		until,
		currentTime,
		currentTime,
		"Do smth",
		status,
	}

	if taskMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", taskMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(ID.Bytes())

	mock.ExpectQuery("INSERT INTO tasks").WithArgs(UserID, "TaskOne",
		until, currentTime, currentTime, "Do smth").WillReturnRows(rows)

	if _, err := AddTask(task); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTask(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, taskMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-00000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")
	var status bool

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")

	expected := Task{
		ID:        ID,
		UserID:    UserID,
		Name:      "TaskOne",
		Time:      until,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Desc:      "Do smth",
		Completed: status,
	}

	if taskMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", taskMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Time", "CreatedAt", "UpdatedAt", "Desc", "Completed"}).
		AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime, "Do smth", status)

	mock.ExpectQuery("SELECT (.+) FROM tasks").WithArgs(ID).WillReturnRows(rows)

	result, err := GetTask(ID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	}

}

func TestDeleteTask(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, taskMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if taskMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", taskMockErr)
	}

	mock.ExpectExec("DELETE FROM tasks").WithArgs(
		ID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteTask(ID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetTasks(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, taskMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	var status bool

	ID, _ := uuid.FromString("00000000-0000-0000-0000-00000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	var expects = []Task{
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
			status,
		},
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
			status,
		},
	}

	if taskMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", taskMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Time", "CreatedAt", "UpdatedAt", "Desc", "Completed"}).
		AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime,
			"Do smth", status).AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime, "Do smth", status)

	mock.ExpectQuery("SELECT (.+) FROM tasks").WillReturnRows(rows)

	result, err := GetTasks()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i] {
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}

func TestGetUserTasks(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, taskMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	var status bool

	ID, _ := uuid.FromString("00000000-0000-0000-0000-00000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	var expects = []Task{
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
			status,
		},
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
			status,
		},
	}

	if taskMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", taskMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Time", "CreatedAt", "UpdatedAt", "Desc", "Completed"}).
		AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime,
			"Do smth", status).AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime, "Do smth", status)

	mock.ExpectQuery("SELECT (.+) FROM tasks").WithArgs(UserID).WillReturnRows(rows)

	result, err := GetUserTasks(UserID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i] {
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}
