package gonhentai_test

import (
	"github.com/KiritoNya/gonhentai"
	"testing"
)

func TestNewUser(t *testing.T) {
	_, err := gonhentai.NewUser(InputTests.UserUrl)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUser_GetId(t *testing.T) {

	// Create user object
	u, err := gonhentai.NewUser(InputTests.UserUrl)
	if err != nil {
		t.Fatal(err)
	}

	// Get id
	err = u.GetId()
	if err != nil {
		t.Fatal(err)
	}

	// Check
	if u.Id != OutputTest.User.Id {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.User.Id, u.Id)
	}
}
