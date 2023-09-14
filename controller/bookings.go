package controller

import (
	"miniproject/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooking godoc
// @Summary Get user's bookings
// @Description Get a list of user's bookings
// @Tags Bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.Bookings
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /booking [get]
func (cn *Controllers) GetBooking(ctx *gin.Context) {
	userId, ok := ctx.Get("loggedInUser")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid token, missing 'user' in context"))
		return
	}

	userlogin, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid 'user id' type in context"))
		return
	}

	bookings, err := cn.Controller.FindBooking(userlogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(500, "failed to retrieve bookings"))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse("200 - ok", bookings))
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking for the user
// @Tags Bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body model.Bookings true "Booking details"
// @Success 201 {object} model.Bookings
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /booking [post]
func (cn *Controllers) CreateBooking(ctx *gin.Context) {
	userId, ok := ctx.Get("loggedInUser")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid token, missing 'user' in context"))
		return
	}

	userlogin, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid 'user id' type in context"))
		return
	}

	var reqbookings model.Bookings

	if err := ctx.ShouldBindJSON(&reqbookings); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "invalid input "+err.Error()))
		return
	}

	reqbookings.Status = "check in"

	booking, err := cn.Controller.Addboking(userlogin, reqbookings.Room_id, reqbookings.Total_day, reqbookings.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "booking failed"))
		return
	}

	ctx.JSON(http.StatusCreated, NewResponse("201 - success booking", booking))

}

// UpdateBooking godoc
// @Summary Update a booking
// @Description Update a booking for the user
// @Tags Bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Booking ID"
// @Param input body model.Bookings true "Updated booking details"
// @Success 201 {object} model.Bookings
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /booking/{id} [put]
func (cn *Controllers) UpdateBooking(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	userId, ok := ctx.Get("loggedInUser")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid token, missing 'user' in context"))
		return
	}

	userlogin, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid 'user' type in context"))
		return
	}

	var reqbookings model.Bookings
	reqbookings.Status = "checkin"

	if err := ctx.ShouldBindJSON(&reqbookings); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "invalid input "+err.Error()))
		return
	}

	booking, err := cn.Controller.EditBooking(intId, userlogin, reqbookings.Room_id, reqbookings.Total_day, reqbookings.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "booking failed"))
		return
	}

	ctx.JSON(http.StatusCreated, NewResponse("201 - success to edit booking", booking))

}
