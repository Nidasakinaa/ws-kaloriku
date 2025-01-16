package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"` 
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

type Customer struct {
	ID        	 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      	 string             `bson:"name,omitempty" json:"name,omitempty"`
	Phone     	 string             `bson:"phone,omitempty" json:"phone,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type ReqMenu struct {
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"` 
	Image       string             `bson:"image,omitempty" json:"image,omitempty"`
}

type ReqCustomer struct {
	Name      	 string             `bson:"name,omitempty" json:"name,omitempty"`
	Phone     	 string             `bson:"phone,omitempty" json:"phone,omitempty"`
}

type ReqAdmin struct {
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}