package routes

import (
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/handlers"
	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo, db *config.DB) {

	api := e.Group("/api/v1.1.0")
	//User
	userHandler := handlers.NewUser(db)
	userAPIs := api.Group("/users")
	userAPIs.POST("/", userHandler.HandleCreateUserRequest())
	userAPIs.GET("/", userHandler.HandleGetUserListRequest())
	//Gist
	gistHandler := handlers.NewGist(db)
	gistAPIs := api.Group("/gists")
	gistAPIs.GET("", gistHandler.HandleGetGistListByUserIDRequest())
	gistAPIs.GET("/:referenceid", gistHandler.HandleGistByReferenceIDRequest())
	gistAPIs.POST("/refresh", gistHandler.HandleGistRefresherRequest())

}
