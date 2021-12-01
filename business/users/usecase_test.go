package users_test

import (
	"context"
	"errors"
	"miniproject/app/middleware"
	"miniproject/business/users"
	"miniproject/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = mocks.Repository{Mock: mock.Mock{}}
var userService users.UseCase
var token string
var userDomain users.Domain
var AllUserDomain []users.Domain

func setup() {
	userService = users.NewUseCase(&userRepository, time.Hour*10, &middleware.ConfigJWT{})
	userDomain = users.Domain{
		Id: 1,
		Name: "agung",
		Email: "agung@gmail.com",
		Address: "lampung",
		Password: "bujangorgen",
	}
	token = "tokennnnn"
	AllUserDomain = append(AllUserDomain, userDomain)
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil)
	userRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(users.Domain{}, nil)
	t.Run("Test Case 1 | Success Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), users.Domain{
			Id: 1,
			Name: "agung",
			Email: "agung@gmail.com",
			Address: "lampung",
			Password: "bujangorgen",
		})

		assert.NoError(t, err)
		assert.Equal(t, userDomain, user)
	})

	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, errors.New("Email Empty"))
	t.Run("Test Case 2 | Error Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), users.Domain{
			Id: 1,
			Name: "agung",
			Email: "",
			Address: "lampung",
			Password: "2323232",
		})

		assert.Error(t, err)
		assert.NotNil(t, user)
	})

	userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, errors.New("Password Empty"))
	t.Run("Test Case 3 | Error Register", func(t *testing.T) {
		user, err := userService.Register(context.Background(), users.Domain{
			Id: 1,
			Name: "agung",
			Email: "dsadsadas@gmail.com",
			Address: "lampung",
			Password: "",
		})

		assert.Error(t, err)
		assert.NotNil(t, user)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Success Login", func(t *testing.T) {
		setup()
		userRepository.On("GetByEmail",
			mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		user, token, err := userService.Login(context.Background(), userDomain.Email, "123")

		assert.NotNil(t, token)
		assert.NoError(t, err)
		assert.Equal(t, user, users.Domain{})
	})

	t.Run("Test Case 2 | Cannot Login (Password Not Found)", func(t *testing.T) {
		data, token, err := userService.Login(context.Background(), userDomain.Email, "")

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("Test Case 3 | Cannot Login (Email Not Found)", func(t *testing.T) {
		data, token, err := userService.Login(context.Background(), "", userDomain.Password)

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)
		assert.Equal(t, token, "")
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("Test case 1 | Success GetAllUsers", func(t *testing.T){
		setup()
		userRepository.On("GetAllUsers", mock.Anything, mock.Anything).Return(AllUserDomain, nil).Once()
		user, err := userService.GetAllUsers(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userDomain.Name, user[0].Name)
	})

	t.Run("Test case 2 | Error GetAllUsers", func(t *testing.T) {
		setup()
		userRepository.On("GetAllUsers", mock.Anything, mock.Anything).Return([]users.Domain{}, errors.New("Users Not Found")).Once()
		data, err := userService.GetAllUsers(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []users.Domain{})
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test case 1 | Success GetUserById", func(t *testing.T) {
		setup()
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetUserById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	t.Run("Test case 2 | Error GetUserById(user Id = 0)", func(t *testing.T) {
		setup()
		userDomain.Id = 0
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.GetUserById(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, users.Domain{})
	})

	t.Run("Test case 3 | Error GetUserById", func(t *testing.T) {
		setup()
		userRepository.On("GetUserById", mock.Anything, mock.AnythingOfType("uint")).Return(users.Domain{}, nil).Once()
		data, err := userService.GetUserById(context.Background(), 7)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, data, users.Domain{})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test case 1 | Success Update", func(t *testing.T) {
		setup()
		userRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, nil).Once()
		data, err := userService.Update(context.Background(), userDomain, userDomain.Id)

		assert.NotNil(t, data)
		assert.NoError(t, err)
	})

	t.Run("Test case 2 | Failed Update", func(t *testing.T) {
		setup()
		userRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userDomain, errors.New("Users Not Found")).Once()
		data, err := userService.Update(context.Background(), userDomain, userDomain.Id)

		assert.Equal(t, data, users.Domain{})
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test case 1 | Success Delete", func(t *testing.T) {
		setup()
		userRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		err := userService.Delete(context.Background(), userDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("Test case 2 | Failed Delete", func(t *testing.T) {
		setup()
		userRepository.On("Delete", mock.Anything, mock.Anything).Return(errors.New("Users Not Found")).Once()
		err := userService.Delete(context.Background(), userDomain.Id)

		assert.Equal(t, err, errors.New("Users Not Found"))
		assert.Error(t, err)
	})
}