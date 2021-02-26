package github

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
	config.LoadConfig(os.Getenv(common.EnvURL))
	//#################################################
	handler := NewGitHub()
	t.Run("getPublicGists", func(t *testing.T) {
		res, err := handler.GetPublicGistsByUsername("a-berahman")
		if err != nil {
			t.Fatalf("expected error to be nil ")
		}
		if res == nil {
			t.Fatalf("expected res to be not nil ")
		}
		if res[0].Owner.Login != "a-berahman" {
			t.Fatalf("expected username to be a-berahman ")
		}

	})

}
