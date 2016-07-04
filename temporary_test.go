package temp

import (
	"os"
	"testing"
)

type Session struct {
	ID string
	T
}

var (
	testSession = &Session{
		T: T{},
	}
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
