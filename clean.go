package temp

import (
	"fmt"
	"reflect"
	"time"
)

//Clean blocks forever and removes expires keys from a map with values that adhere to Temporary.
//Clean panics if m is not a map or slice.
//Clean makes heavy use of reflection.
func Clean(collection interface{}, interval time.Duration) {
	val := reflect.ValueOf(collection)

	if !val.Type().Elem().Implements(reflect.TypeOf((*Temporary)(nil)).Elem()) {
		panic("Map does not have elements that implement Temporary")
	}

	var elem Temporary

	switch val.Kind().String() {
	case "map":
		for {
			for _, key := range val.MapKeys() {
				elem = val.MapIndex(key).Interface().(Temporary)
				if Expired(elem) {
					val.SetMapIndex(key, reflect.Value{})
				}
				time.Sleep(interval)
			}
			time.Sleep(time.Second)
		}
	default:
		panic(fmt.Sprintf("Clean expects a map but got: %v", val))

	}
}
