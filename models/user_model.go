package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Age       int                `json:"age" bson:"age"`
	Address   string             `json:"address" bson:"address"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func (u *User) Create() {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
}
