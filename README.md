# Package `geval` - Documentation

## Overview

The `geval` package in Go provides a simple utility to handle the evaluation and execution of tasks with the ability to manage context, error handling, and verbose output. It allows developers to define custom actions for success, failure, and panic scenarios. The core components of the package are the `Params` and `Context` types.

## Why use `geval`?

The `geval` package offers several advantages that make it useful for various types of software development:

1. **Context Management**: The package provides a `Context` type, allowing developers to manage and pass contextual information across different functions and goroutines. This facilitates sharing data and communication among concurrent tasks.

2. **Error Handling**: It allows users to define custom error handlers, which can be used to perform specific actions based on the success or failure of a task. This helps in handling errors gracefully and taking appropriate actions.

3. **Panic Handling**: Developers can specify a panic handler to gracefully handle unexpected panics and recover from them. This ensures that a panic in one goroutine doesn't crash the entire application.

4. **Verbose Output**: The package supports verbose output, making it easier to debug and understand the execution flow. It provides developers with more insights into the tasks' progress and status.

5. **Concurrency Support**: The package is designed to be concurrency-friendly, enabling developers to create and manage goroutines sharing the same context. This makes it easier to execute tasks concurrently without race conditions.

## Examples

### Basic Use:

#### Example 1: Simple Task Execution

```go
package main

import (
	"context"
	"fmt"
	"github.com/jaavier/geval"
)

func main() {
	ctx := geval.CreateContext()

	params := &geval.Params{
		Context: ctx,
		Handler: simpleTask,
	}

	geval.Run(params)
}

func simpleTask(ctx *geval.Context) error {
	// Perform the main task here
	fmt.Println("Task executed successfully!")
	return nil
}
```

### Intermediate Use:

#### Example 2: Error Handling and Context Management

```go
package main

import (
	"context"
	"fmt"
	"github.com/jaavier/geval"
)

func main() {
	ctx := geval.CreateContext()
	ctx.Update("user_id", 123)

	params := &geval.Params{
		Context: ctx,
		Handler: errorProneTask,
		Failed:  handleFailure,
	}

	geval.Run(params)
}

func errorProneTask(ctx *geval.Context) error {
	// Simulate an error condition
	return fmt.Errorf("something went wrong")
}

func handleFailure(ctx *geval.Context) {
	// Custom error handling
	fmt.Println("Task failed!")
	userId := ctx.Read("user_id")
	fmt.Printf("User ID: %v\n", userId)
}
```

### Advanced Use:

#### Example 3: Verbose Output

```go
package main

import (
	"context"
	"fmt"
	"github.com/jaavier/geval"
)

func main() {
	ctx := geval.CreateContext()

	params := &geval.Params{
		Context: ctx,
		Handler: verboseTask,
		Success: handleSuccess,
		Verbose: true,
	}

	geval.Run(params)
}

func verboseTask(ctx *geval.Context) error {
	// Some task execution
	return nil
}

func handleSuccess(ctx *geval.Context) {
	// Custom success handler with verbose output
	fmt.Println("Task executed successfully!")
}
```

#### Example 4: Concurrent Task Execution with Shared Context and Channel

```go
package main

import (
	"fmt"
	"sync"

	"github.com/jaavier/geval"
)

func main() {
	ctx := geval.CreateContext()

	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ctx, &wg)
	go consumer(ctx, &wg)

	wg.Wait()
}

func producer(ctx *geval.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate producing data and sending it through the channel
	for i := 1; i <= 5; i++ {
		ctx.Channel <- i
		fmt.Println("Producer sent:", i)
	}
	close(ctx.Channel)
}

func consumer(ctx *geval.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate consuming data from the channel
	for data := range ctx.Channel {
		fmt.Println("Consumer received:", data)
	}
}
```
