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

func (r *mutationResolver) CreateTransaction(ctx context.Context, input models.TransactionInput) (*models.Transaction, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateTransaction(ctx context.Context, input models.TransactionInput) (*models.Transaction, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id int) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Transactions(ctx context.Context, id *int) ([]*models.Transaction, error) {
	panic("not implemented")
}
