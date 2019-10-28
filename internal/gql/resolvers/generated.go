package resolvers

import (
  "context"
  // "log"

	"github.com/romainm/buddy-server/internal/orm"
	dbm "github.com/romainm/buddy-server/internal/orm/models"
	"github.com/romainm/buddy-server/internal/gql"
	"github.com/romainm/buddy-server/internal/gql/models"
	tf "github.com/romainm/buddy-server/internal/gql/resolvers/transformations"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	ORM *orm.ORM
}

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
  // entity := "transactions"
  whereID := "id = ?"
  records := []*models.Transaction{}
  dbRecords := []*dbm.Transaction{}
  db := r.ORM.DB.New()
  if id != nil {
    db = db.Where(whereID, *id)
  }
  db = db.Find(&dbRecords)
  for _, dbRec := range dbRecords {
    if rec, err := tf.TransactionToGraph(dbRec); err != nil {
      // log.Print(entity, err)
    } else {
      records = append(records, rec)
    }
  }
  return records, db.Error
}
