package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
	IsActive  bool               `bson:"isActive" json:"isActive"`
	CreatedOn time.Time          `bson:"createdOn" json:"createdOn"`
	UpdatedOn time.Time          `bson:"updatedOn" json:"updatedOn"`
}

// Initialize sets default values for BaseModel.
func (bm *BaseModel) Initialize() {
	now := time.Now().UTC()
	bm.IsDeleted = false
	bm.IsActive = true
	bm.CreatedOn = now
	bm.UpdatedOn = now
}

// UpdateOn sets the UpdatedOn field to the current time.
func (bm *BaseModel) UpdateOn() {
	bm.UpdatedOn = time.Now().UTC()
}
