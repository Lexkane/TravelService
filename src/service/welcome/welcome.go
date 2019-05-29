package welcome

import (
	"TravelService/src/service/common"
	"net/http"
)

type welcomeStruct struct {
	Message string `json:"message"`
}

// GetWelcomeHandler get policies for partner
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderJSON(w, r, welcomeStruct{Message: "Hello World"})

}
