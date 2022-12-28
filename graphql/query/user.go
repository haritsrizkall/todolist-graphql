package query

import (
	"context"
	"todolist-graphql/graphql/resolver"
	"todolist-graphql/service"
)

type UserQuery struct {
	UserService service.UserService
}

func NewUserQuery(service service.UserService) *UserQuery {
	return &UserQuery{
		UserService: service,
	}
}

func (q *UserQuery) Users(ctx context.Context) (*[]*resolver.UserResolver, error) {
	users, err := q.UserService.FindAll()
	if err != nil {
		return nil, err
	}
	var userResolvers []*resolver.UserResolver
	for _, user := range users {
		userResolvers = append(userResolvers, &resolver.UserResolver{User: user})
	}
	return &userResolvers, nil
}

type UserArgs struct {
	ID int32
}

func (q *UserQuery) User(ctx context.Context, args UserArgs) (*resolver.UserResolver, error) {
	user, err := q.UserService.FindById(args.ID)
	if err != nil {
		return nil, err
	}
	return &resolver.UserResolver{User: user}, nil
}
