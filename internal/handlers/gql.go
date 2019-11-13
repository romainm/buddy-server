package handlers

import (
    "context"
    "net/http"
    "io/ioutil"

    "github.com/romainm/buddy-server/internal/orm"
    // "github.com/romainm/buddy-server/internal/gql"
    // "github.com/romainm/buddy-server/internal/gql/resolvers"
    "github.com/gin-gonic/gin"
    graphql "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
    // // NewExecutableSchema and Config are in the generated.go file
    // c := gql.Config{
    //     Resolvers: &resolvers.Resolver{
	// 		ORM: orm,
	// 	},
    // }

    // h := handler.GraphQL(gql.NewExecutableSchema(c))

    // return func(c *gin.Context) {
    //     h.ServeHTTP(c.Writer, c.Request)
    // }
    return func(c *gin.Context) {
        c.String(http.StatusOK, "OK")
    }
}

type query struct{}

func (_ *query) Transactions(ctx context.Context, args struct{ ID *graphql.ID }) string { 
    return "Hello, world!" 
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
    s, err := getSchema("internal/gql/schemas/schema.graphql")
    if err != nil {
        panic(err)
    }
    schema := graphql.MustParseSchema(s, &query{})

    h := relay.Handler{Schema: schema}
    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}
