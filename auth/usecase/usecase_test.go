package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zhashkevych/go-clean-architecture/auth/repository/mock"
	"github.com/zhashkevych/go-clean-architecture/models"
	"testing"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)

	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)

	var (
		username = "user"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Password: "xyz", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
