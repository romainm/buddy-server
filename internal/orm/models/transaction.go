package models

import (
  "github.com/jinzhu/gorm"
)

type Transaction struct {
  gorm.Model
  Name         string
  Amount       float64 
}