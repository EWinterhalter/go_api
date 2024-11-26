package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@test.com",
		Password: "123456",
	}
}
