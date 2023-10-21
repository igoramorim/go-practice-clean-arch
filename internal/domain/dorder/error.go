package dorder

import "errors"

var (
	ErrInvalidID = errors.New("product id is invalid")

	ErrInvalidPrice      = errors.New("product price is invalid")
	ErrInvalidTax        = errors.New("product tax is invalid")
	ErrInvalidFinalPrice = errors.New("product final price is invalid")
)
