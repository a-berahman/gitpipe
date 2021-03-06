package services

import (
	"github.com/a-berahman/gitpipe/common/providers"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/services/github"
	"github.com/a-berahman/gitpipe/services/pipedrive"
)

// GetService returns service instace
// - solution is implemented by Factory design pattern
func GetService(serviceConst int) interface{} {
	switch serviceConst {
	case providers.Pipedrive:
		return pipedrive.NewPipedrive()
	case providers.GitHub:
		return github.NewGitHub()
	}
	return nil
}

// Githuber is implemented by objects that promote GitHub API features
type Githuber interface {
	GetPublicGistsByUsername(username string) (result []models.Gists, err error)
}

// Pipedriver is implemented by objects that promote Pipedrive API features
type Pipedriver interface {
	InsertActivityByGists(gist models.Gists) (activityID int, err error)
	GetActivityByID(id int) (models.Activity, error)
}
