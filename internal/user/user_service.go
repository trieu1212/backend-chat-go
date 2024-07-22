package user

import (
	"context"
	"fmt"
	"golang-test/utils"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateUser(ctx context.Context, req CreateUserReq) (CreateUserRes, error) {
	hassPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return CreateUserRes{}, err
	}

	u := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hassPass,
	}

	user, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return CreateUserRes{}, err
	}

	return CreateUserRes{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s service) Login(ctx context.Context, req LoginReq) (LoginRes, error) {
	fmt.Println(req)
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginRes{}, err
	}
	
	err = utils.CheckPassword(user.Password, req.Password)
	if err != nil {
		return LoginRes{}, err
	}

	token, err := utils.CreateToken(user.ID.Hex(), user.Role)
	if err != nil {
		return LoginRes{}, err
	}

	return LoginRes{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}
