package gist

import (
	"context"
	"fmt"
	"time"

	"github.com/a-berahman/gitpipe/config"
	"github.com/a-berahman/gitpipe/models"
	"github.com/a-berahman/gitpipe/utility/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

//GistRepository presents basic property for Gist repository
type GistRepository struct {
	db  *config.DB
	log *zap.SugaredLogger
}

//NewGist makes new instance of Gist Repository
func NewGist(db *config.DB) *GistRepository {
	return &GistRepository{log: logger.Logger(), db: db}
}

//Create creates new gist
func (g *GistRepository) Create(title, userID, referenceID string) error {
	//Here we initializing gist model for inserting in DB
	gistModel := models.Gist{}
	gistModel.ID = primitive.NewObjectID()
	gistModel.UserID = userID
	gistModel.ReferenceID = referenceID
	gistModel.Title = title
	gistModel.CreateAt = time.Now()

	res, err := g.db.Gist.InsertOne(context.Background(), gistModel) //inserting gist in db
	if err != nil {
		g.log.Errorw("failed to insert gist in db",
			"gist model", gistModel,
			"error", err,
		)
		return err
	}
	//check the result of insertOne method
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		g.log.Errorw("failed to convert insert ID",
			"InsertID", res.InsertedID,
		)
		return fmt.Errorf("failed to convert insert ID")
	}
	g.log.Infow("inserted new gist successfully",
		"gist model", gistModel,
		"insertID", oid.Hex(),
	)

	return nil

}

// GetByUserID returns gist by userID
func (g *GistRepository) GetByUserID(userID string) ([]*models.Gist, error) {

	// Here's an array in which you can store the decoded documents
	var result []*models.Gist
	filter := bson.M{"user_id": userID}
	curr, err := g.db.Gist.Find(context.TODO(), filter)
	// Close the cursor once finished
	defer curr.Close(context.TODO())
	if err != nil {
		g.log.Errorw("failed to get gists by user",
			"error", err,
		)
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for curr.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var gist models.Gist
		err := curr.Decode(&gist)
		if err != nil {
			g.log.Errorw("failed to decode Gist model",
				"error", err,
				"gist", curr,
			)
			continue
		}
		result = append(result, &gist)
	}

	return result, nil
}
