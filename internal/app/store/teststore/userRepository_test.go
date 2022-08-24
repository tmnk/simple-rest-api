package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmnk/simple-rest-api/internal/app/model"
	"github.com/tmnk/simple-rest-api/internal/app/store"
	"github.com/tmnk/simple-rest-api/internal/app/store/teststore"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}

func TestUserRepository_FindId(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	s.User().Create(u)
	_, err := s.User().FindId(u.ID + 1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	_, err = s.User().FindId(u.ID)
	assert.NoError(t, err)
}
