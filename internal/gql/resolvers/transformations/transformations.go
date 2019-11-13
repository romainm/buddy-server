package transformations

import (
  gql "github.com/romainm/buddy-server/internal/gql/models"
  dbm "github.com/romainm/buddy-server/internal/orm/models"
)

func TransactionToGraph(i *dbm.Transaction) (o *gql.Transaction, err error) {
  id_ := int(i.ID)
  o = &gql.Transaction{
    ID:          id_,
    Name:        i.Name,
    Amount:      i.Amount,
  }
  return o, err
}

func TransactionInputToDb(i *gql.TransactionInput) (o *dbm.Transaction, err error) {
  o = &dbm.Transaction{
    Name:        i.Name,
    Amount:      i.Amount,
  }
  return o, err
}
