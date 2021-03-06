package model_test

import (
	"TravelService/src/database"
	"TravelService/src/model"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var (
	mock    sqlmock.Sqlmock
	mockErr error
)

func TestSaveTrainToTrip(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	trainID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	mock.ExpectExec("INSERT INTO trips_trains").WithArgs(trainID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := model.AddToTrip(trainID, tripID, model.Train{}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
