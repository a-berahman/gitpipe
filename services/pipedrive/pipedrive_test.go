package pipedrive

import (
	"fmt"
	"testing"

	"github.com/a-berahman/gitpipe/common"
	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
)

func TestFirst(t *testing.T) {
	config.LoadConfig(fmt.Sprintf("%v%v", common.RootDir, "env.yaml"))
	//#################################################
	handler := NewPipedrive()

	t.Run("InsertActivityByGists", func(t *testing.T) {
		activityID, err := handler.InsertActivityByGists(models.Gists{
			ID:         "3ffa4fd7653165df9b103bf722f47289",
			ExteraInfo: `{"url":"https://api.github.com/gists/3ffa4fd7653165df9b103bf722f47289","forks_url":"https://api.github.com/gists/3ffa4fd7653165df9b103bf722f47289/forks","commits_url":"https://api.github.com/gists/3ffa4fd7653165df9b103bf722f47289/commits","id":"3ffa4fd7653165df9b103bf722f47289","node_id":"MDQ6R2lzdDNmZmE0ZmQ3NjUzMTY1ZGY5YjEwM2JmNzIyZjQ3Mjg5","git_pull_url":"https://gist.github.com/3ffa4fd7653165df9b103bf722f47289.git","git_push_url":"https://gist.github.com/3ffa4fd7653165df9b103bf722f47289.git","html_url":"https://gist.github.com/3ffa4fd7653165df9b103bf722f47289","files":{"Create Graph":{"filename":"Create Graph","type":"text/plain","language":null,"raw_url":"https://gist.githubusercontent.com/a-berahman/3ffa4fd7653165df9b103bf722f47289/raw/e6c441304de49db8e93d263856a6030955ac1fcb/Create%20Graph","size":839}},"public":true,"created_at":"2020-07-17T07:01:42Z","updated_at":"2020-07-17T07:01:42Z","description":"","comments":0,"user":null,"comments_url":"https://api.github.com/gists/3ffa4fd7653165df9b103bf722f47289/comments","owner":{"login":"a-berahman","id":62759025,"node_id":"MDQ6VXNlcjYyNzU5MDI1","avatar_url":"https://avatars.githubusercontent.com/u/62759025?v=4","gravatar_id":"","url":"https://api.github.com/users/a-berahman","html_url":"https://github.com/a-berahman","followers_url":"https://api.github.com/users/a-berahman/followers","following_url":"https://api.github.com/users/a-berahman/following{/other_user}","gists_url":"https://api.github.com/users/a-berahman/gists{/gist_id}","starred_url":"https://api.github.com/users/a-berahman/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/a-berahman/subscriptions","organizations_url":"https://api.github.com/users/a-berahman/orgs","repos_url":"https://api.github.com/users/a-berahman/repos","events_url":"https://api.github.com/users/a-berahman/events{/privacy}","received_events_url":"https://api.github.com/users/a-berahman/received_events","type":"User","site_admin":false},"truncated":false}`,
		})
		if err != nil {
			t.Fatalf("expected error to be nil ")
		}

		if activityID == 0 {
			t.Fatalf("expected activityID to be non-zero")
		}
	})
}
