package handlers

import (
	"net/http"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/common/repositorytype"
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/repository"
	"github.com/a-berahman/gitpipe/utility/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserHandler is a HttpHandler that presents basic property for user
type UserHandler struct {
	log            *zap.SugaredLogger
	UserRepository repository.Userer
}

// NewUser creates user handler instance
func NewUser(db *config.DB) *UserHandler {
	return &UserHandler{log: logger.Logger(), UserRepository: repository.GetRepository(repositorytype.User, db).(repository.Userer)}
}

type createUserRequest struct {
	Username string `json:"username"`
}

//HandleCreateUserRequest prepares process of creating user request
func (u *UserHandler) HandleCreateUserRequest() func(c echo.Context) error {
	u.log.Infow("calling the CreateUser")
	return func(c echo.Context) error {
		req := new(createUserRequest)
		err := c.Bind(req)
		if err != nil {
			return echo.ErrBadRequest
		}

		err = u.createUser(req.Username)
		if err != nil {
			u.log.Errorw("failed to create user ",
				"error", err,
				"username", req.Username,
			)
			return c.JSON(http.StatusNotFound, common.H{"message": "OK"})
		}

		return c.JSON(http.StatusOK, common.H{"message": "OK"})
	}

}

// HandleGetUserListRequest prepares process of getting user list request
func (u *UserHandler) HandleGetUserListRequest() func(c echo.Context) error {
	u.log.Infow("calling the GetUserList")
	return func(c echo.Context) error {

		res, err := u.getUserList()
		if err != nil {
			u.log.Errorw("failed to get user list ",
				"error", err,
				"result", res,
			)
			return c.JSON(http.StatusNotFound, common.H{"message": "OK"})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (u *UserHandler) createUser(username string) error {
	err := u.UserRepository.Create(username)
	if err != nil {
		u.log.Errorw("failed to handle create user process by username",
			"error", err,
			"username", username,
		)
		return err
	}
	return nil
}

func (u *UserHandler) getUserList() ([]*models.User, error) {
	res, err := u.UserRepository.GetAll()
	if err != nil {
		u.log.Errorw("failed to handle create user process by username",
			"error", err,
			"result", res,
		)
		return nil, err
	}
	return res, nil
}
