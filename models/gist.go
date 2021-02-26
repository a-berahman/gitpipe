package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Gist represents Gist data model
type Gist struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id"`
	ReferenceID int                `json:"reference_id" bson:"reference_id"`
	Title       string             `json:"title" bson:"title"`
	CreateAt    time.Time          `json:"create_at" bson:"create_at"`
}
