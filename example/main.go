package main

import (
	"fmt"
	"time"

	"github.com/ammario/temp"
)

type session struct {
	ID string
	temp.T
}

func main() {
	sess := session{}
	temp.ExpireAfter(&sess, time.Second)
	fmt.Printf("Session expired: %v\n", temp.Expired(&sess))
	time.Sleep(time.Second)
	fmt.Printf("Session expired: %v\n", temp.Expired(&sess))
}
