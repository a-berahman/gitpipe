package user

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/config"
)

func TestFirstUser(t *testing.T) {
	if common.AppMode == common.TestMode {
		os.Setenv("ENV_URL", common.ConfigDir)
	}
	db := config.LoadConfig(os.Getenv(common.EnvURL))
	//#################################################
	handler := NewUser(db)
	username := fmt.Sprintf("%v_%v", "user", time.Now().Format(time.Kitchen))
	t.Run("CreateUserSuccessfully", func(t *testing.T) {
		err := handler.Create(username)
		if err != nil {
			t.Fatal("expected error to be nil")
		}
	})

	t.Run("GetUserByUsername", func(t *testing.T) {
		got, err := handler.GetByUsername(username)
		if err != nil {
			t.Fatal("expected error to be nil")
		}
		if got.Username != username {
			t.Fatalf("expected username to be %v \n got: %v \n", username, got.Username)
		}
	})

	t.Run("GetAllUser", func(t *testing.T) {
		got, err := handler.GetAll()
		if err != nil {
			t.Fatal("expected error to be nil")
		}
		if len(got) < 1 {
			t.Fatal("expected users length to be more than 0")
		}
	})

	t.Run("UpdateUserByUsername", func(t *testing.T) {
		err := handler.UpdateLastCheck(username)
		if err != nil {
			t.Fatal("expected error to be nil")
		}

	})

}
