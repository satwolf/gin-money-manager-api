package mocks

import (
	"gin-money-manager-api/modules/shared/repository/options"
	"gin-money-manager-api/modules/user/entity"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByUsername(username string) (*entity.User, error) {
	args := m.Called(username)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Create(user entity.User) (entity.User, error) {
	args := m.Called(user)

	result, _ := args.Get(0).(entity.User)
	return result, args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) Find(id string, options *options.FindOptions) (entity.User, error) {
	args := m.Called(id, options)

	result, _ := args.Get(0).(entity.User)
	return result, args.Error(1)
}

func (m *MockUserRepository) FindAll(options *options.FindAllOptions) ([]entity.User, int, error) {
	args := m.Called(options)

	users, _ := args.Get(0).([]entity.User)
	count, _ := args.Get(1).(int)

	return users, count, args.Error(2)
}

func (m *MockUserRepository) Update(user entity.User) (entity.User, error) {
	args := m.Called(user)

	result, _ := args.Get(0).(entity.User)
	return result, args.Error(1)
}

func (m *MockUserRepository) FindBy(where map[string]any) ([]entity.User, error) {
	args := m.Called(where)

	users, _ := args.Get(0).([]entity.User)
	return users, args.Error(1)
}

func (m *MockUserRepository) DB() *gorm.DB {
	args := m.Called()

	db, _ := args.Get(0).(*gorm.DB)
	return db
}
