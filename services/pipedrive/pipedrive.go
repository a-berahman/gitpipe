package pipedrive

import (
	"fmt"
	"time"

	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/utility/logger"
	"go.uber.org/zap"
)

//Pipedrive is a type of Pipedrive API
type Pipedrive struct {
	log *zap.SugaredLogger
}

//NewPipedrive makes new instance of GitHub type
func NewPipedrive() *Pipedrive {
	return &Pipedrive{
		log: logger.Logger(),
	}
}

// InsertActivityByGists adds new activity via pipedrive api
func (p *Pipedrive) InsertActivityByGists(gist models.Gists) (activityID int, err error) {
	fmt.Println(time.Now().Format("2006-01-02"))
	req := AddActivityRQ{
		DueDate: time.Now().Format("2006-01-02"),
		Note:    gist.ExteraInfo,
		Subject: gist.ID,
	}
	res := AddActivityRS{}
	if err = sendPostRequestAndCheckResponse(req, &res, fmt.Sprintf("%v%v", config.CFG.Pipedrive.MainURL, config.CFG.Pipedrive.AddActivityURL)); err != nil {
		p.log.Errorw("failed to send post request",
			"URL", fmt.Sprintf("%v%v", config.CFG.Pipedrive.MainURL, config.CFG.Pipedrive.AddActivityURL),
			"error", err,
		)
		return 0, err

	}
	if !res.Success {
		p.log.Errorw("insertActivityByGist has some problem in response",
			"Success", res.Success,
			"Error", res.Error.Error,
			"ErrorCode", res.Error.ErrorCode,
		)
		return 0, fmt.Errorf(res.Error.Error)
	}
	return res.Data.ID, nil
}

//GetActivityByID returns an activity by ID
func (p *Pipedrive) GetActivityByID(id int) (activity models.Activity, err error) {

	res := GetActivityRS{}
	if err = sendGetRequestAndCheckResponse(&res, fmt.Sprintf("%v%v", config.CFG.Pipedrive.MainURL, fmt.Sprintf(config.CFG.Pipedrive.GetActivityURL, id)), make(map[string]string)); err != nil {
		p.log.Errorw("failed to send get activity request",
			"URL", fmt.Sprintf("%v%v", config.CFG.Pipedrive.MainURL, fmt.Sprintf(config.CFG.Pipedrive.GetActivityURL, id)),
			"error", err,
		)
		return activity, err

	}
	if !res.Success {
		p.log.Errorw("GetActivityByID has some problem in response",
			"Success", res.Success,
			"Error", res.Error.Error,
			"ErrorCode", res.Error.ErrorCode,
		)

		return activity, fmt.Errorf(res.Error.Error)
	}

	activity.Note = res.Data.Note
	activity.ID = res.Data.ID
	activity.Subject = res.Data.Subject

	return activity, nil
}
