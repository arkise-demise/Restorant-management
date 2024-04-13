package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID		primitive.ObjectID `bson:"_id"`
	Name		*string `bson:"_id"`
	Price		*float64 `json:"_id"`
	Food_image		
	Created_at		
	Updated_at		
	Food_id		string `json:"food_id"`
	Menu_id		*string				`json:"menu_id" validate:"required"`
}