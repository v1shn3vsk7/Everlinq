package models

type User struct {
	Id       string `json:"id"       bson:"_id, omitempty"`
	UserName string `json:"username" bson:"username"`
	Password string	`json:"password" bson:"password"`
	Email    string	`json:"email"    bson:"email"`
}
