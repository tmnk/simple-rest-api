package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmnk/simple-rest-api/internal/app/model"
	"github.com/tmnk/simple-rest-api/internal/app/store"
	"github.com/tmnk/simple-rest-api/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}
