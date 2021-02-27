package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/common/repositorytype"
	"github.com/a-berahman/gitpipe/common/servicetype"
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/repository"
	"github.com/a-berahman/gitpipe/services"
	"github.com/a-berahman/gitpipe/utility/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// GistHandler is a HttpHandler that presents basic property for gist
type GistHandler struct {
	log            *zap.SugaredLogger
	GistRepository repository.Gister
	UserRepository repository.Userer
	PipeService    services.Pipedriver
	GitService     services.Githuber
}

// NewGist creates gist handler instance
func NewGist(db *config.DB) *GistHandler {
	return &GistHandler{log: logger.Logger(),
		GistRepository: repository.GetRepository(repositorytype.Gist, db).(repository.Gister),
		UserRepository: repository.GetRepository(repositorytype.User, db).(repository.Userer),
		PipeService:    services.GetService(servicetype.Pipedrive).(services.Pipedriver),
		GitService:     services.GetService(servicetype.GitHub).(services.Githuber),
	}
}

//HandleGetGistListByUserIDRequest prepares process of getting gist list request
func (g *GistHandler) HandleGetGistListByUserIDRequest() func(c echo.Context) error {
	g.log.Infow("calling the GetGistListByUserID handler")
	return func(c echo.Context) error {
		userID := c.QueryParam("userid")
		res, err := g.getGistListByUserID(userID)
		if err != nil {
			g.log.Errorw("failed to get gist list by userid ",
				"error", err,
				"user id", userID,
			)
			return c.JSON(http.StatusNotFound, common.H{"message": "OK"})
		}

		return c.JSON(http.StatusOK, res)
	}

}

// HandleGistByReferenceIDRequest prepares process of getting gist request
func (g *GistHandler) HandleGistByReferenceIDRequest() func(c echo.Context) error {
	g.log.Infow("calling the GistByReferenceID")
	return func(c echo.Context) error {
		referenceID, _ := strconv.Atoi(c.Param("referenceid"))

		res, err := g.getGistByReferenceID(referenceID)
		if err != nil {
			g.log.Errorw("failed to get gist by reference id ",
				"error", err,
				"result", res,
			)
			return c.JSON(http.StatusNotFound, common.H{"message": "OK"})
		}

		return c.JSON(http.StatusOK, res)
	}
}

//HandleGistRefresherRequest prepares process of refreshing gist request
func (g *GistHandler) HandleGistRefresherRequest() func(c echo.Context) error {
	g.log.Infow("calling the refresh gist handler")
	return func(c echo.Context) error {

		err := g.gistRefresher()
		if err != nil {
			g.log.Errorw("failed to refresh gist",
				"error", err,
			)
			return c.JSON(http.StatusNotFound, common.H{"message": "OK"})
		}

		return c.JSON(http.StatusOK, common.H{"message": "OK"})
	}

}

//############################

func (g *GistHandler) getGistListByUserID(userID string) (map[string]*models.Gist, error) {
	g.log.Infow("calling the GetGistListByUserID")
	res, err := g.GistRepository.GetByUserID(userID)
	if err != nil {
		g.log.Errorw("failed to get gist list by userID ",
			"error", err,
			"user id", userID,
			"result", res,
		)
		return nil, err
	}
	return res, nil
}

func (g *GistHandler) getGistByReferenceID(referenceID int) (string, error) {
	g.log.Infow("calling the GetGistByReferenceID")
	res, err := g.PipeService.GetActivityByID(referenceID)
	if err != nil {
		g.log.Errorw("failed to get gist list by userID ",
			"error", err,
			"reference ID", referenceID,
			"result", res,
		)
		return "", err
	}

	return res.Note, nil
}

func (g *GistHandler) gistRefresher() error {
	g.log.Infow("calling the gistRefresher handler")
	users, err := g.UserRepository.GetAll()
	if err != nil {
		g.log.Errorw("failed to get user list",
			"error", err,
		)
		return err
	}
	//traversing in user list and Excute gistrefresh action via Future/Promise design pattern
	for _, user := range users {
		currOBJ := &DO{}
		var wg sync.WaitGroup
		wg.Add(1)
		currOBJ.Success(func(username string) {
			g.log.Infow("refresh succeed",
				"username", username,
			)
			wg.Done()
		}).Fail(func(username string, err error) {
			g.log.Infow("refresh failed",
				"username", username,
				"error", err,
			)
			wg.Done()
		})

		currOBJ.Execute(gistRefreshContext(g, *user))

	}
	return nil
}

//gistRefreshContext returns ExcuteFunction for Executing
func gistRefreshContext(g *GistHandler, user models.User) ExecuteFunc {
	return func() (string, error) {
		return gistRefreshAction(g, user)
	}
}
func gistRefreshAction(g *GistHandler, user models.User) (username string, err error) {

	currGistList, err := g.GistRepository.GetByUserID(user.ID.Hex())
	if err != nil {
		g.log.Errorw("failed to get gist list from DB by UserID",
			"error", err,
			"user id", user.ID.Hex(),
		)
	}

	gitGistList, err := g.GitService.GetPublicGistsByUsername(user.Username)
	if err != nil {
		g.log.Errorw("failed to get public gist by username",
			"error", err,
			"username", user.Username,
		)
	}
	for _, v := range gitGistList {
		if _, ok := currGistList[v.ID]; ok {
			continue
		}
		referenceID, err := g.PipeService.InsertActivityByGists(v)
		if err != nil {
			g.log.Errorw("failed to insert gist as an activity ",
				"error", err,
				"gist", v,
			)
		}
		g.GistRepository.Create(v.ID, user.ID.Hex(), referenceID)
	}
	return user.Username, err
}
