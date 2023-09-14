package controller

import (
	"miniproject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAvailableRoom godoc
// @Summary Find available rooms
// @Description Find available rooms for booking
// @Tags Rooms
// @Accept json
// @Produce json
// @Success 200 {array} model.Rooms
// @Failure 404 {object} ErrorResponse
// @Router /rooms [get]
func (cn *Controllers) FindAvailableRoom(ctx *gin.Context) {
	rooms := []model.Rooms{}
	result, err := cn.Controller.FindAvailableRooms(rooms)
	if err != nil {
		ctx.JSON(http.StatusNotFound, NewErrorResponse(404, "room not found"))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse("200 ok", result))
}
