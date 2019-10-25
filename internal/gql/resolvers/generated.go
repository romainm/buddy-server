package resolvers

import (
	"context"

	"github.com/romainm/buddy-server/internal/gql"
	"github.com/romainm/buddy-server/internal/gql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	id_ :=     "ec17af15-e354-440c-a09f-69715fc8b595"
	email :=  "your@email.com"
	userId := "UserID-1"
	 records := []*models.User{
        &models.User{
            ID:     &id_,
            Email:  &email,
            UserID: &userId,
        },
    }
    return records, nil
}
