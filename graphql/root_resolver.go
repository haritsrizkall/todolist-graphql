package graphql

import (
	"todolist-graphql/graphql/mutation"
	"todolist-graphql/graphql/query"
)

type RootResolver struct {
	*query.UserQuery
	*mutation.UserMutation
}

func NewRootResolver(
	userQuery *query.UserQuery,
	userMutation *mutation.UserMutation,
) *RootResolver {
	return &RootResolver{
		userQuery,
		userMutation,
	}
}
