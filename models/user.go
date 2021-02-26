package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User represents User data model
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	LastCheck time.Time          `json:"last_check" bson:"last_check"`
	CreateAt  time.Time          `json:"create_at" bson:"create_at"`
}
