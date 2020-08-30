package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUserInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(r)

	if value, ok := pathParams["userID"]; ok {
		userID, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("[ERR][getUserInfo] Error while convreting user-id : %s to int, Err : %v", value, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(messsageInvalidUserID))
			return
		}
	}

	w.Write([]byte(fmt.Sprintf(`{"message": %d"}`, userID)))
	w.WriteHeader(http.StatusOK)
	return
}
