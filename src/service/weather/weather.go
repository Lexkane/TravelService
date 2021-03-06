package weather

import (
	"TravelService/src/model"
	"TravelService/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

func GetWeatherByTrainIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	trainID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trainID", err)
		return
	}

	result, err := model.GetWeatherByTrainID(trainID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: This train is not connected with your trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
func GetWeatherByFlightIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	flightID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong flightID", err)
		return
	}

	result, err := model.GetWeatherByFlightID(flightID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: This flight is not connected with your trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
