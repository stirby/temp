# Temp
Temporary structs and maps in Golang

# Basic Usage
## Temporary struct

```go
type session struct {
	ID string
	temp.T
}

func main() {
	sess := session{}
	temp.ExpireAfter(&sess, time.Second)
	fmt.Printf("Session expired: %v\n", temp.Expired(&sess)) // false
	time.Sleep(time.Second)
	fmt.Printf("Session expired: %v\n", temp.Expired(&sess)) // true
}

```


## Temporary maps

```go
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
go Clean(m, time.Millisecond*50) //Clean blocks forever
time.Sleep(time.Second * 2)
//Map should be empty here
```

