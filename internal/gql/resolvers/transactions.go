package resolvers

import (
  "context"
  // "log"

	dbm "github.com/romainm/buddy-server/internal/orm/models"
	"github.com/romainm/buddy-server/internal/gql/models"
	tf "github.com/romainm/buddy-server/internal/gql/resolvers/transformations"
)

func (r *mutationResolver) CreateTransaction(ctx context.Context, input models.TransactionInput) (*models.Transaction, error) {
  dbo, err := tf.TransactionInputToDb(&input)
  if err != nil {
    return nil, err
  }
  // Create scoped clean db interface
  db := r.ORM.DB.New().Begin()
  db = db.Create(dbo).First(dbo) // Create the user
  gql, err := tf.TransactionToGraph(dbo)
  if err != nil {
    db.RollbackUnlessCommitted()
    return nil, err
  }
  db = db.Commit()
  return gql, db.Error
}

func (r *mutationResolver) UpdateTransaction(ctx context.Context, input models.TransactionInput) (*models.Transaction, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTransaction(ctx context.Context, id int) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Transactions(ctx context.Context, id* int) ([]*models.Transaction, error) {
  // entity := "transactions"
  whereID := "id = ?"
  records := []*models.Transaction{}
  dbRecords := []*dbm.Transaction{}
  db := r.ORM.DB.New()
  if id != nil {
    db = db.Where(whereID, id)
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
