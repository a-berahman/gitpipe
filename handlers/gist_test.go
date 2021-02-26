package handlers

import (
	"os"
	"testing"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/config"
)

func TestFirst(t *testing.T) {
	if common.AppMode == common.TestMode {
		os.Setenv("ENV_URL", common.ConfigDir)
	}
	db := config.LoadConfig(os.Getenv(common.EnvURL))
	//#################################################
	handler := NewGist(db)

	t.Run("GistRefresher", func(t *testing.T) {
		err := handler.gistRefresher()
		if err != nil {
			t.Fatalf("expected error to bi nil, error: %v", err)
		}
	})
}
