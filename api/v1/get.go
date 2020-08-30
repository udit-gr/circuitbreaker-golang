package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUserInfo return info for the user-id in request payload
func GetUserInfo(w http.ResponseWriter, r *http.Request) {

	var (
		userID int
		err    error
	)

	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	if value, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(value)
		if err != nil {
			log.Printf("[ERR][getUserInfo] Error while converting user-id : %s to int, Err : %v", value, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(messsageInvalidUserID))
			return
		}
	}

	w.Write([]byte(fmt.Sprintf(`{"message": "%d"}`, userID)))
	return
}
