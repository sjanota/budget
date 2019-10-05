package storage

import (
	"errors"

	"github.com/sjanota/budget/backend/pkg/storage/errorcode"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoBudget         = errors.New("budget does not exist")
	ErrAlreadyExists    = errors.New("already exists")
	ErrDoesNotExists    = errors.New("does not exist")
	ErrInvalidReference = errors.New("invalid reference to other resource")
)

func isDuplicateKeyError(err error) bool {
	writeException, ok := err.(mongo.WriteException)
	if !ok {
		return false
	}

	for _, writeError := range writeException.WriteErrors {
		if writeError.Code == errorcode.DuplicateKey {
			return true
		}
	}
	return false
}
