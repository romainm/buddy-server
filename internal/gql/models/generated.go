// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Transaction struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TransactionInput struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}
