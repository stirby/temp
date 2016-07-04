package temp

import "time"

//Temporary defines the necessary methods used by an expiring object
type Temporary interface {
	Expires() time.Time
	SetExpires(ti time.Time)
}

//ExpireAfter sets a T to expire after a duration from now
func ExpireAfter(t Temporary, dur time.Duration) {
	t.SetExpires(time.Now().Add(dur))
}

//Expired checks if a Temporary has expired
func Expired(t Temporary) bool {
	return time.Now().After(t.Expires())
}
