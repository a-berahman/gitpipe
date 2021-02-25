package user

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

//UserRepository presents basic property for user repository
type UserRepository struct {
	db  *config.DB
	log *zap.SugaredLogger
}

//NewUser makes new instance of User Repository
func NewUser(db *config.DB) *UserRepository {
	return &UserRepository{log: logger.Logger(), db: db}
}

//Create creates new user with username
func (u *UserRepository) Create(username string) error {
	//initialize user model
	userModel := models.User{}
	userModel.ID = primitive.NewObjectID()
	userModel.CreateAt = time.Now()
	userModel.Username = username
	//insert user
	res, err := u.db.User.InsertOne(context.Background(), userModel)
	if err != nil {
		u.log.Errorw("failed to insert user in db",
			"user model", userModel,
			"error", err,
		)
		return err
	}
	//check the result of insertOne method
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		u.log.Errorw("failed to convert insert ID",
			"InsertID", res.InsertedID,
		)
		return fmt.Errorf("failed to convert insert ID")
	}
	u.log.Infow("inserted new user successfully",
		"user model", userModel,
		"insertID", oid.Hex(),
	)

	return nil
}

// GetByUsername returns user by username
func (u *UserRepository) GetByUsername(username string) (models.User, error) {
	currentUser := models.User{}
	filter := bson.M{"username": username}
	res := u.db.User.FindOne(context.Background(), filter)
	if err := res.Decode(&currentUser); err != nil {
		u.log.Errorw("failed to decode user model",
			"error", err,
		)
		return currentUser, err
	}

	return currentUser, nil
}

// GetAll returns list of user
func (u *UserRepository) GetAll() ([]*models.User, error) {

	// Here's an array in which you can store the decoded documents
	var result []*models.User

	cur, err := u.db.User.Find(context.TODO(), bson.M{})
	// Close the cursor once finished
	defer cur.Close(context.TODO())
	if err != nil {
		u.log.Errorw("failed to get user list",
			"error", err,
		)
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			u.log.Errorw("failed to decode user model",
				"error", err,
				"user", cur,
			)
			continue
		}
		result = append(result, &user)
	}

	return result, nil
}
