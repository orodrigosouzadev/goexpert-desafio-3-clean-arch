package entity

import "errors"

const ErrInvalidID = "invalid id"
const ErrInvalidPrice = "invalid price"
const ErrInvalidTax = "invalid tax"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:         id,
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New(ErrInvalidID)
	}
	if o.Price <= 0 {
		return errors.New(ErrInvalidPrice)
	}
	if o.Tax <= 0 {
		return errors.New(ErrInvalidTax)
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
