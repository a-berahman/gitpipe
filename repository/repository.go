package repository

import (
	"github.com/a-berahman/gitpipe/common/repositorytype"
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/repository/gist"
	"github.com/a-berahman/gitpipe/repository/user"
)

// GetRepository returns repository  instace
// - solution is implemented by Factory design pattern
func GetRepository(repositoryConst int, db *config.DB) interface{} {
	switch repositoryConst {
	case repositorytype.Gist:
		return gist.NewGist(db)
	case repositorytype.User:
		return user.NewUser(db)
	}
	return nil
}

//Gister is implemented by objects that promote Gist Repository
type Gister interface {
	Create(title, userID string, referenceID int) error
	GetByUserID(userID string) (map[string]*models.Gist, error)
}

//Userer is implemented by objects that promote User  Repository
type Userer interface {
	Create(username string) error
	GetByUsername(username string) (models.User, error)
	GetAll() ([]*models.User, error)
	UpdateLastCheck(username string) error
}
