package transformations

import (
  gql "github.com/romainm/buddy-server/internal/gql/models"
  dbm "github.com/romainm/buddy-server/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func TransactionToGraph(i *dbm.Transaction) (o *gql.Transaction, err error) {
  id_ := int(i.ID)
  o = &gql.Transaction{
    ID:          id_,
    Name:        i.Name,
    Amount:      i.Amount,
  }
  return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func TransactionInputToDb(i *gql.TransactionInput) (o *dbm.Transaction, err error) {
  o = &dbm.Transaction{
    Name:        i.Name,
    Amount:      i.Amount,
  }
  return o, err
}
