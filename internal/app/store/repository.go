package store

import "github.com/tmnk/simple-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindId(int) (*model.User, error)
}
