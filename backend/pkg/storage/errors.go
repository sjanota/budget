package storage

import "errors"

var (
	ErrNoBudget              = errors.New("budget does not exist")
	ErrAccountAlreadyExists  = errors.New("account already exists")
	ErrAccountDoesNotExists  = errors.New("account does not exist")
	ErrEnvelopeAlreadyExists = errors.New("envelope already exists")
	ErrEnvelopeDoesNotExists = errors.New("envelope does not exist")
	ErrCategoryAlreadyExists = errors.New("category already exists")
)
