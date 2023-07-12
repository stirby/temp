package temp

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

//Clean blocks forever and removes expires keys from a map with values that adhere to Temporary.
//Clean panics if m is not a map.
//Clean makes heavy use of reflection.
//scanInterval is slept between whole scans of a map, if zero it will block the thread.
//checkInterval is slept between checks of individual elements.
func Clean(m interface{}, mutex *sync.RWMutex, scanInterval time.Duration, checkInterval time.Duration) {
	val := reflect.ValueOf(m)

	if !val.Type().Elem().Implements(reflect.TypeOf((*Temporary)(nil)).Elem()) {
		panic("Map does not have elements that implement Temporary")
	}

	var elem Temporary

	switch val.Kind().String() {
	case "map":
		for {

			for _, key := range val.MapKeys() {
				mutex.Lock()
				elem = val.MapIndex(key).Interface().(Temporary)
				if Expired(elem) {
					val.SetMapIndex(key, reflect.Value{})
				}
				mutex.Unlock()
				time.Sleep(checkInterval)
			}

			time.Sleep(scanInterval)
		}
	default:
		panic(fmt.Sprintf("Clean expects a map but got: %v", val))

	}
}
