package internal

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccessToken(t *testing.T) {
	user := &domain.User{
		ID:   "123",
		Name: "testuser",
	}
	secret := "mysecret"
	expiry := 1 // 1 hour
	token, err := CreateAccessToken(user, secret, expiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestIsAuthorized(t *testing.T) {
	user := &domain.User{
		ID:   "123",
		Name: "testuser",
	}
	secret := "mysecret"
	expiry := 1 // 1 hour

	token, err := CreateAccessToken(user, secret, expiry)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	isAuth, err := IsAuthorized(token, secret)
	assert.NoError(t, err)
	assert.True(t, isAuth)
}
