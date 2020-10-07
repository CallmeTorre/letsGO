package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CallmeTorre/letsGO/mvc/services"
	"github.com/CallmeTorre/letsGO/mvc/utils"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userID, err := (strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		userError := utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(userError)
		response.WriteHeader(userError.StatusCode)
		response.Write(jsonValue)
	}
	user, apiError := services.GetUser(userID)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
	}
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
