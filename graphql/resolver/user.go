package resolver

import "todolist-graphql/entity"

type UserResolver struct {
	User *entity.User
}

func (r *UserResolver) ID() int32 {
	return r.User.ID
}

func (r *UserResolver) Username() string {
	return r.User.Username
}
