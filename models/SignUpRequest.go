package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignUp struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username" form:"Username" binding:"required" bson:"username"`
	Password string             `json:"password" form:"Password" binding:"required" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
}
