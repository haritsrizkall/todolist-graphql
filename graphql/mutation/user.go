package mutation

import (
	"context"
	"todolist-graphql/entity"
	"todolist-graphql/graphql/resolver"
	"todolist-graphql/service"
)

type UserMutation struct {
	UserService service.UserService
}

func NewUserMutation(service service.UserService) *UserMutation {
	return &UserMutation{
		UserService: service,
	}
}

type CreateUserArgs struct {
	Username string
}

func (m *UserMutation) CreateUser(ctx context.Context, args CreateUserArgs) (*resolver.UserResolver, error) {
	user, err := m.UserService.Create(&entity.User{
		Username: args.Username,
	})
	if err != nil {
		return nil, err
	}

	return &resolver.UserResolver{User: user}, nil
}
