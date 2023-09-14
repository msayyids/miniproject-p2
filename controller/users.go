package controller

import (
	"fmt"
	"miniproject/helper"
	"miniproject/model"
	"miniproject/repo"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"

	_ "miniproject/docs"
)

type Controllers struct {
	Controller repo.Repo
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param User body model.Users true "User registration details"
// @Success 201 {object} Response
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func (cn *Controllers) Register(ctx *gin.Context) {
	reqBodyUser := model.Users{}

	reqBodyUser.Deposit_amount = 0

	if err := ctx.ShouldBindJSON(&reqBodyUser); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "invalid input "+err.Error()))
		return
	}

	NewUsers, err := cn.Controller.AddUsers(reqBodyUser.Name, reqBodyUser.Email, helper.HashedPassword(reqBodyUser.Password), reqBodyUser.PhoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "failed to register"))
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "sekedot@gmail.com")
	m.SetHeader("To", NewUsers.Email)
	m.SetHeader("Subject", "Welcome to Hocation!")

	emailBody := `
        <html>
        <body>
            <p>Dear ` + NewUsers.Name + `,</p>
            <p>Welcome to Hocation! You have successfully registered with us.</p>
            <p>We are excited wait for your coming</p>
            <p>Best regards,</p>
            <p>The Hocation Team</p>
        </body>
        </html>
    `
	m.SetBody("text/html", emailBody)

	emailPassword := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer("smtp.elasticemail.com", 2525, "sekedot@gmail.com", emailPassword)

	if err := d.DialAndSend(m); err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(500, "failed to send registration email"))
		return
	}

	ctx.JSON(http.StatusCreated, NewResponse("201 - status created", "Success! You have registered with Hocation, "+NewUsers.Name+"!"))
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param email body string true "User email"
// @Param password body string true "User password"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func (cn *Controllers) Login(ctx *gin.Context) {
	var reqBodyUser struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBodyUser); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "invalid input "+err.Error()))
		return
	}

	userLogin, err := cn.Controller.FindUserByEmail(reqBodyUser.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "failed to login,invalid email or password"))
		return
	}

	isValidPassword := helper.ValidatePasword(userLogin.Password, reqBodyUser.Password)
	if !isValidPassword {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "failed to login,invalid email or password"))
		return
	}

	token, err := helper.GenerateToken(userLogin.Id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "failed to generate token"))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse("200 - success login", token))

}

// EditAmount godoc
// @Summary Edit user's deposit amount
// @Description Edit user's deposit amount
// @Tags Users
// @Accept json
// @Produce json
// @Param amount body int true "New deposit amount"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /topup [put]
func (cn *Controllers) EditAmount(ctx *gin.Context) {

	userId, ok := ctx.Get("loggedInUser")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid token, missing 'user' in context"))
		return
	}

	var reqBody struct {
		Amount int `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(400, "invalid input "+err.Error()))
		return
	}

	userLogin, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid 'user id' type in context"))
		return
	}

	err := cn.Controller.EditAmount(reqBody.Amount, userLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(500, fmt.Sprintf("Failed to edit amount: %v", err)))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse("200 - success", fmt.Sprintf("Deposit amount updated  %d", reqBody.Amount)))
}

// GetLoggedInUserInfo godoc
// @Summary Get information of the logged-in user
// @Description Get information of the user who is currently logged in
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Users
// @Failure 401 {object} ErrorResponse
// @Router /user [get]
func (cn *Controllers) GetLoggedInUserInfo(ctx *gin.Context) {
	userId, ok := ctx.Get("loggedInUser")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid token, missing 'user' in context"))
		return
	}

	userLogin, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, NewErrorResponse(401, "Invalid 'user id' type in context"))
		return
	}

	userInfo, err := cn.Controller.FindById(userLogin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(500, "failed to retrieve user information"))
		return
	}

	ctx.JSON(http.StatusOK, NewResponse("200 - ok", userInfo))
}
