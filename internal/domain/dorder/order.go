package dorder

func New(id int64, price, tax float64) (*Order, error) {
	orderID, err := NewID(id)
	if err != nil {
		return nil, err
	}

	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	if tax <= 0 {
		return nil, ErrInvalidTax
	}

	return &Order{
		id:         orderID,
		price:      price,
		tax:        tax,
		finalPrice: 0,
	}, nil
}

type Order struct {
	id         ID
	price      float64
	tax        float64
	finalPrice float64
}

func (o *Order) CalculateFinalPrice() error {
	finalPrice := o.price + o.tax
	if finalPrice <= 0 {
		return ErrInvalidFinalPrice
	}

	o.finalPrice = finalPrice
	return nil
}
