package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	Role     string             `json:"role" bson:"role,omitempty"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type Service interface {
	CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error)
	Login(ctx context.Context, req LoginReq) (LoginRes, error)
}

// Create
type CreateUserReq struct {
	Name     string `json:"name" required:"true"`
	Email    string `json:"email", required:"true"`
	Password string `json:"password", required:"true"`
}

type CreateUserRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Login
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}


