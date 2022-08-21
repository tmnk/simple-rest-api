package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "us@sss.com",
		Password: "45545454545",
	}
}
