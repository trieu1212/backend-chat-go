package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db mongo.Database
}

func NewRepository(db mongo.Database) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateUser(ctx context.Context, user User) (User, error) {
	col := r.db.Collection("users")

	u := User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     "user",
	}

	_, err := col.InsertOne(ctx, u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	col := r.db.Collection("users")

	var user User
	err := col.FindOne(ctx, User{Email: email}).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
