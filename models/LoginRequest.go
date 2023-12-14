package models

type LoginRequest struct {
	Username string `json:"username" form:"Username" binding:"required" bson:"username"`
	Password string `json:"password" form:"Password" binding:"required" bson:"password"`
}
