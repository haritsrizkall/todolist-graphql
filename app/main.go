package main

import (
	"context"
	"fmt"
	"time"
	"todolist-graphql/drivers/mysql"
	"todolist-graphql/entity"
	our_graphql "todolist-graphql/graphql"
	"todolist-graphql/graphql/mutation"
	"todolist-graphql/graphql/query"
	"todolist-graphql/graphql/schema"
	"todolist-graphql/repository"
	"todolist-graphql/service"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&entity.Todo{},
		&entity.User{},
	)
}

type Todo struct {
	ID        int32
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type rootResolver struct {
	TodoQuery
}

func newRootResolver(
	todoQuery TodoQuery,
) *rootResolver {
	return &rootResolver{
		TodoQuery: todoQuery,
	}
}

type TodoQuery struct{}

func NewTodoQuery() *TodoQuery {
	return &TodoQuery{}
}

func (q *TodoQuery) Todos(ctx context.Context) (*[]*TodoResolver, error) {
	todos := []Todo{
		{
			ID:        1,
			Title:     "Todo 1",
			Status:    "Active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Todo 2",
			Status:    "Active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	var todoResolvers []*TodoResolver
	for _, todo := range todos {
		todoResolvers = append(todoResolvers, &TodoResolver{Data: todo})
	}
	return &todoResolvers, nil
}

type GetByIdArgs struct {
	ID int32
}

func (q *TodoQuery) Todo(ctx context.Context, args GetByIdArgs) (*TodoResolver, error) {
	todo := Todo{
		ID:        1,
		Title:     "Todo 1",
		Status:    "Active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &TodoResolver{Data: todo}, nil
}

type TodoResolver struct {
	Data Todo
}

func (r *TodoResolver) ID() int32 {
	return r.Data.ID
}

func (r *TodoResolver) Title() string {
	return r.Data.Title
}

func (r *TodoResolver) Status() string {
	return r.Data.Status
}

func (r *TodoResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.Data.CreatedAt}
}

func (r *TodoResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.Data.UpdatedAt}
}

func main() {
	fmt.Println("Hello World")
	dbConfig := mysql.DbConfig{
		User:     "root",
		Password: "florist123",
		Host:     "127.0.0.1",
		Port:     "3306",
		Name:     "todolist",
	}
	db := dbConfig.InitDb()
	dbMigrate(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	e := echo.New()
	userQuery := query.NewUserQuery(userService)
	userMutation := mutation.NewUserMutation(userService)
	rootResolver := our_graphql.NewRootResolver(
		userQuery,
		userMutation,
	)
	schemaStr, err := schema.String()
	if err != nil {
		fmt.Println(err)
	}
	schema := graphql.MustParseSchema(schemaStr, rootResolver)
	e.POST("/graphql", echo.WrapHandler(&relay.Handler{Schema: schema}))
	e.File("/graphiql", "web/graphiql.html")
	e.Logger.Fatal(e.Start(":1323"))
}
