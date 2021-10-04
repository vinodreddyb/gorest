package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"restexample/models"
	"restexample/repository"
)

var useRepository repository.UserRepository

func init() {
	useRepository = repository.NewUserRepository()
}

func AddNewUser(c *fiber.Ctx) error {
	user := &models.User{}
	err := c.BodyParser(user)

	if err != nil {
		errorResp := prepareErrorResponse(err, http.StatusNotAcceptable, "Body parsing error")
		return c.Status(errorResp.Code).JSONP(errorResp)
	}
	id, err := useRepository.Save(*user)

	if err != nil {
		errorResp := prepareErrorResponse(err, http.StatusInternalServerError, "Internal Server error")
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	user.ID = id
	resp := prepareSuccessResponse(user)
	return c.Status(resp.Code).JSON(resp)

}

func GetAllUsers(c *fiber.Ctx) error {
	users := useRepository.GetAllUsers()
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    users,
		Title:   "OK",
		Message: "All users",
	}

	return c.Status(resp.Code).JSON(resp)
}

func prepareSuccessResponse(user *models.User) models.Response {
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    user,
		Title:   "OK",
		Message: "new user added successfully",
	}
	return resp
}

func prepareErrorResponse(err error, code int, message string) models.Response {

	errorResp := models.Response{
		Code:    code,
		Body:    err.Error(),
		Title:   "Error",
		Message: message,
	}
	return errorResp
}
