package models

type Employee struct {
	ID        string `json:"uuid" bson:"uuid"`
	Name      string             `json:"name" bson:"name"`
	UserName  string             `json:"username" bson:"username"`
	Password  string 			 `json:"-" bson:"password"`
	Department string            `json:"department" bson:"department"`
}
