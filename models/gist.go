package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Gist presents gist schema for DB
type Gist struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      string             `json:"user_id" bson:"user_id"`
	ReferenceID string             `json:"reference_id" bson:"reference_id"`
	Title       string             `json:"username" bson:"username"`
	CreateAt    time.Time          `json:"create_at" bson:"create_at"`
}
