// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package database

import ()

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}
