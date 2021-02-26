package gist

import (
	"os"
	"testing"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/config"
	"github.com/google/uuid"
)

func TestFirstGist(t *testing.T) {
	if common.AppMode == common.TestMode {
		os.Setenv("ENV_URL", common.ConfigDir)
	}

	db := config.LoadConfig(os.Getenv(common.EnvURL))
	//#################################################
	handler := NewGist(db)
	userid := uuid.New().String()
	t.Run("GetGistByUserID", func(t *testing.T) {

		err := handler.Create(uuid.New().String(), userid, 0)
		if err != nil {
			t.Fatal("expected error to be nil")
		}
	})

	t.Run("GetGistByUserID", func(t *testing.T) {
		got, err := handler.GetByUserID(userid)
		if err != nil {
			t.Fatal("expected error to be nil")
		}
		if len(got) < 1 {
			t.Fatal("expected got length to be more than 0")
		}
	})
}
