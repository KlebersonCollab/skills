package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// User represents a user in the system.
type User struct {
	ID   string
	Name string
}

// UserRepository defines the interface for user persistence.
type UserRepository interface {
	GetUser(id string) (*User, error)
}

// MockUserRepository is a mock implementation of UserRepository.
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(id string) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

// UserService handles business logic for users.
type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUserName(id string) (string, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func TestUserService_GetUserName(t *testing.T) {
	// Setup mock.
	mockRepo := new(MockUserRepository)
	service := &UserService{repo: mockRepo}

	// Define expectation.
	mockUser := &User{ID: "123", Name: "Gopher"}
	mockRepo.On("GetUser", "123").Return(mockUser, nil)

	// Execute.
	name, err := service.GetUserName("123")

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, "Gopher", name)
	mockRepo.AssertExpectations(t)
}
