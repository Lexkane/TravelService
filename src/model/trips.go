package model

import (
	"TravelService/src/database"
	"github.com/satori/go.uuid"
	"log"
	"reflect"
)

const (
	createTrip        = "INSERT INTO trips (user_id) VALUES ($1) RETURNING trip_id;"
	getTripIDByUserID = "SELECT trips.trip_id FROM trips WHERE trips.user_id = $1;"
	getTripsByTripID  = "SELECT trips.user_id FROM trips WHERE trip_id = $1;"
)

//Trip is a representation of Event Trip in DB
type Trip struct {
	TripID      uuid.UUID
	UserID      uuid.UUID
	Events      interface{}
	Flights     interface{}
	Museums     interface{}
	Restaurants interface{}
	Hotels      interface{}
	Trains      interface{}
	TotalSum    int
}

//AddTrip creates Trip and saves it to DB
var AddTrip = func(trip Trip) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(createTrip, trip.UserID).Scan(&id)

	return id, err
}

//GetTrip gets Trips from DB by tripID
var GetTrip = func(id uuid.UUID) (Trip, error) {

	var (
		trip Trip
		err  error
	)

	trip.TripID = id

	trip.Events, err = GetFromTrip(id, Event{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	trip.TotalSum += getSum(trip.Events)

	trip.Flights, err = GetFromTrip(id, Flight{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	trip.TotalSum += getSum(trip.Flights)

	trip.Museums, err = GetFromTrip(id, Museum{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	trip.TotalSum += getSum(trip.Museums)

	trip.Hotels, err = GetFromTrip(id, Hotel{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	trip.TotalSum += getSum(trip.Hotels)

	trip.Trains, err = GetFromTrip(id, Train{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	trip.TotalSum += getSum(trip.Trains)

	trip.Restaurants, err = GetFromTrip(id, Restaurant{})
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}

	errDB := database.DB.QueryRow(getTripsByTripID, id).Scan(&trip.UserID)
	if err != nil {
		log.Println(errDB)
		return Trip{}, err
	}

	return trip, nil
}

//GetTripIDsByUserID gets Trips from DB by userID
var GetTripIDsByUserID = func(id uuid.UUID) ([]uuid.UUID, error) {

	var (
		tripIDs []uuid.UUID
		tripID  uuid.UUID
	)

	rows, err := database.DB.Query(getTripIDByUserID, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&tripID); err != nil {
			return nil, err
		}
		tripIDs = append(tripIDs, tripID)
	}
	return tripIDs, nil
}

func getSum(obj interface{}) (totalSum int) {

	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(obj)
		for i := 0; i < s.Len(); i++ {
			totalSum += int(s.Index(i).FieldByName("Price").Int())
		}
	}
	return
}
