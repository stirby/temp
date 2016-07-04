//Package temp provides the ability to have temporary structs
//and simple cleaning of slices and maps containing temporary data
package temp

import "time"

//T is meant to be inherited by instances intended to expire
type T struct {
	expires time.Time
}

//Expires returns when T expires
func (t *T) Expires() time.Time {
	return t.expires
}

//SetExpires sets the expire time for a T
func (t *T) SetExpires(ti time.Time) {
	t.expires = ti
}
