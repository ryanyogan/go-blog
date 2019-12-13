package dbclient

import (
	"github.com/ryanyogan/go-blog/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient --
type MockBoltClient struct {
	mock.Mock
}

// QueryAccount --
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

// OpenBoltDB --
func (m *MockBoltClient) OpenBoltDB() {
	// NO-OP
}

// Seed --
func (m *MockBoltClient) Seed() {
	// NO-OP
}
