package auth

import (
	"fmt"
	"os"
	"testing"
)

/*go test -v test/auth_test.go */
func TestNewJWT(t *testing.T) {
	manager, err := NewAuthManager("secret")
	if err != nil {
		t.Errorf("Auth manager init errors: %s\n", err)
		return
	}

	token, err := manager.NewJWT("15", 1)

	if err != nil {
		t.Errorf("New jwt errors: %s\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Token was generated successfuly: %s\n", token)

	jwtSub, err := manager.ParseJWT(token)

	if err != nil {
		t.Errorf("JWT parse errors: %s \n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Token was parsed successfuly: %s \n", jwtSub)

}
