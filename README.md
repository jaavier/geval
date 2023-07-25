# Start using

Install package: `go get github.com/jaavier/geval`

# Handle Error and Success

```golang
package main

import (
	"fmt"

	"github.com/jaavier/geval"
)

func main() {
	var data []byte
	var err error

	data, err = retrieveData(true)

	geval.Run(&geval.Params{
		Err:     err,
		Success: transformData(data),
	})

	_, err = retrieveData(false)

	geval.Run(&geval.Params{
		Err: err,
		Failed: func() {
			fmt.Println("Catch error:", err.Error())
		},
	})
}

func transformData(data []byte) func() {
	return func() {
		fmt.Println("Bytes raw:", data)
		fmt.Println("Bytes to string:", string(data))
	}
}

func retrieveData(isActive bool) ([]byte, error) {
	if isActive {
		return []byte("hello world"), nil
	}
	return []byte{}, fmt.Errorf("cannot retrieve when isActive = false")
}
```