package services

import (
	"github.com/a-berahman/gitpipe/common/serviceType"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/services/github"
	"github.com/a-berahman/gitpipe/services/pipedrive"
)

// GetService returns service instace
// - solution is implemented by Factory design pattern
func GetService(serviceConst int) interface{} {
	switch serviceConst {
	case serviceType.Pipedrive:
		return pipedrive.NewPipedrive()
	case serviceType.GitHub:
		return github.NewGitHub()
	}
	return nil
}

// Githuber is implemented by objects that promote GitHub API features
type Githuber interface {
	GetPublicGistsByUsername(string) (models.Gists, error)
}

// Pipedriver is implemented by objects that promote Pipedrive API features
type Pipedriver interface {
	InsertActivityByGists(models.Gists) (int, error)
}
