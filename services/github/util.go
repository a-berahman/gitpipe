package github

import (
	"encoding/json"

	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/utility/logger"
	"github.com/a-berahman/gitpipe/utility/rest"
)

func sendGetRequestAndCheckResponse(resModel interface{}, url string, headers map[string]string) error {
	err := makeGetHTTPCall(resModel, url, headers)
	if err != nil {

		logger.Logger().Errorw("http GET call to Jibit failed",
			"error", err,
			"URL", url,
			"RS", resModel,
		)
		return err
	}
	return nil

}

func makeGetHTTPCall(resModel interface{}, urlConst string, headers map[string]string) error {

	res, err := rest.GetJSON(urlConst, headers)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &resModel)
	if err != nil {
		return err
	}

	return nil
}
func initializeRequestBody(req interface{}) ([]byte, error) {
	r, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func addAuthentication(headers map[string]string) {
	headers["Authorization"] = digestAuthrization()
}
func digestAuthrization() string {
	return config.CFG.GitHub.Username + ":" + config.CFG.GitHub.TOKEN
}
