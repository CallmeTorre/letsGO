package controllers

import (
	"net/http"
	"strconv"

	"github.com/CallmeTorre/letsGO/mvc/services"
	"github.com/CallmeTorre/letsGO/mvc/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userError := utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, &userError)
		return
	}
	user, apiError := services.UsersService.GetUser(userID)
	if apiError != nil {
		utils.RespondError(c, apiError)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
