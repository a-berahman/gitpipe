package github

import (
	"encoding/json"
	"fmt"

	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/utility/logger"
	"go.uber.org/zap"
)

//GitHub is a type of GitHub API
type GitHub struct {
	log *zap.SugaredLogger
}

//NewGitHub makes new instance of GitHub type
func NewGitHub() *GitHub {
	return &GitHub{
		log: logger.Logger(),
	}
}

//GetPublicGistsByUsername returns a list of public gits for specific user
func (g *GitHub) GetPublicGistsByUsername(username string) (result []models.Gists, err error) {
	res := []GistsForUserRS{}
	headers := make(map[string]string)
	addAuthentication(headers)
	fmt.Println(username)
	fmt.Println(config.CFG.GitHub.GistURL)
	fmt.Println(config.CFG.GitHub.MainURL)
	fmt.Println(fmt.Sprintf(config.CFG.GitHub.GistURL, username))
	fmt.Println(fmt.Sprintf("%v%v", config.CFG.GitHub.MainURL, fmt.Sprintf(config.CFG.GitHub.GistURL, username)))
	err = sendGetRequestAndCheckResponse(&res,
		fmt.Sprintf("%v%v", config.CFG.GitHub.MainURL, fmt.Sprintf(config.CFG.GitHub.GistURL, username)),
		headers)
	if err != nil {
		g.log.Errorw("failed to send get request",
			"URL", fmt.Sprintf("%v%v", config.CFG.GitHub.MainURL, fmt.Sprintf(config.CFG.GitHub.GistURL, username)),
			"error", err,
		)
		return nil, err
	}

	for _, v := range res {
		jsonNode, _ := json.Marshal(v)
		result = append(result, models.Gists{
			URL:        v.URL,
			ID:         v.ID,
			Owner:      models.Owner(v.Owner),
			ExteraInfo: string(jsonNode),
		})
	}

	return result, nil
}
