package pipedrive

import (
	"encoding/json"
	"fmt"

	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/utility/logger"
	"github.com/a-berahman/gitpipe/utility/rest"
)

func sendPostRequestAndCheckResponse(req interface{}, resModel interface{}, url string) error {

	err := makePostHTTPCall(req, resModel, addToken(url))

	if err != nil {
		logger.Logger().Errorw("http POST call to Pipedrive failed",
			"error", err,
			"RQ", req,
			"RS", resModel,
		)
		return err
	}
	return nil

}
func sendGetRequestAndCheckResponse(resModel interface{}, url string, headers map[string]string) error {
	err := makeGetHTTPCall(resModel, addToken(url), headers)
	if err != nil {

		logger.Logger().Errorw("http GET call to Pipedrive failed",
			"error", err,
			"URL", url,
			"RS", resModel,
		)
		return err
	}
	return nil

}
func makePostHTTPCall(req interface{}, resModel interface{}, url string) error {

	d, err := initializeRequestBody(req)
	if err != nil {
		return err
	}
	res, err := rest.PostHTTPApplicationJSON(url, d)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &resModel)
	if err != nil {
		return err
	}

	return nil
}
func makeGetHTTPCall(resModel interface{}, url string, headers map[string]string) error {

	res, err := rest.GetJSON(url, headers)
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
func addToken(url string) string {
	return fmt.Sprintf("%v?api_token=%v", url, config.CFG.Pipedrive.TOKEN)
}
