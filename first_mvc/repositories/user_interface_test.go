package repositories

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("Not expecting a user with ID = 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("Expecting 404 when user not found")
	}

	if err == nil {
		t.Error("Expecting an error when user ID = 0")
	}
}
