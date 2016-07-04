package temp

import (
	"testing"
	"time"
)

type session struct {
	ID string
	T
}

func TestClean(t *testing.T) {
	m := map[string]*session{
		"123": &session{
			ID: "123",
			T: T{
				expires: time.Now().Add(time.Second),
			},
		},
		"124": &session{
			ID: "124",
			T: T{
				expires: time.Now().Add(time.Second),
			},
		},
		"125": &session{
			ID: "125",
			T: T{
				expires: time.Now().Add(time.Second),
			},
		},
	}
	time.Sleep(time.Second)
	go Clean(m, time.Millisecond*50)
	time.Sleep(time.Second)
	if len(m) > 0 {
		t.Errorf("Map should be completely clean, instead %v elements remain.", len(m))
	}
}
