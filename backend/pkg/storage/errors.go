package storage

import "errors"

var (
	ErrNoBudget         = errors.New("budget does not exist")
	ErrAlreadyExists    = errors.New("already exists")
	ErrDoesNotExists    = errors.New("does not exist")
	ErrInvalidReference = errors.New("invalid reference to other resource")
)
