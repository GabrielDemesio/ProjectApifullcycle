package tests

import (
	"FULLCYCLE/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser2(t *testing.T) {
	user, err := entity.NewUser("Gabriel", "Gabriel@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Gabriel", user.Name)
	assert.Equal(t, "Gabriel@gmail.com", user.Email)
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := entity.NewUser("Gabriel", "Gabriel@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("123456"))
	assert.False(t, user.CheckPassword("12345689"))
	assert.NotEqual(t, "123456", user.Password)
}
