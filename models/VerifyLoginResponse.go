package models

type VerifyLogin struct {
	Username string `json:"username" form:"Username" binding:"required" bson:"username"`
	Password string `json:"password" form:"Password" binding:"required" bson:"password"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
}
