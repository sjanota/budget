package models

import "github.com/pkg/errors"

var (
	ErrMalformedDate  = errors.New("Date must be in format YYYY-MM-DD")
	ErrMalformedMonth = errors.New("Month must be in format YYYY-MM")
)
