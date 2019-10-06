package resolver

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_resolver "github.com/sjanota/budget/backend/pkg/resolver/mocks"
)

func before(t *testing.T) (*Resolver, *mock_resolver.MockStorageMockRecorder, func()) {
	mock := gomock.NewController(t)
	mockStorage := mock_resolver.NewMockStorage(mock)
	resolver := &Resolver{Storage: mockStorage}
	return resolver, mockStorage.EXPECT(), mock.Finish
}
