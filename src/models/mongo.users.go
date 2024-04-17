package models

import "time"

type MongoUsersRepository struct {
	ID            string    `json:"id" bson:"id"`
	PhoneNumber   string    `json:"phone_number" bson:"phone_number"`
	Gender        string    `json:"gender" bson:"gender"`
	FirstName     string    `json:"first_name" bson:"first_name" validate:"required"`
	LastName      string    `json:"last_name" bson:"last_name" validate:"required"`
	Email         string    `json:"email" bson:"email" validate:"required,email"`
	EmailVerified bool      `json:"email_verified" bson:"email_verified"`
	Password      string    `json:"password" bson:"password"`
	ProfilePicURL string    `json:"profile_pic_url" bson:"profile_pic_url"`
	Role          string    `json:"role" bson:"role" default:"geust" validate:"required"`
	Status        string    `json:"status" bson:"status" default:"offline" validate:"required"`
	Token         string    `json:"token" bson:"token" default:"" required:"true"`
	RefreshToken  string    `json:"refresh_token" bson:"refresh_token" default:"" required:"false"`
	CreatedTime   time.Time `json:"created_time" bson:"created_time"`
	UpdatedTime   time.Time `json:"updated_time" bson:"updated_time"`
}
