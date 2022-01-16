package first_test

import (
	"testing"
	"time"
)

/*go test -v test/first_test.go */
func TestFirst(t *testing.T) {

	t.Log(time.Now().Add(time.Hour))
}
