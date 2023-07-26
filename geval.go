package geval

import (
	"context"
	"fmt"
	"runtime"
)

type Params struct {
	Err     error
	Success func(ctx *Context)
	Failed  func(ctx *Context)
	Handler func(ctx *Context) error
	Panic   func(v any) error
	Context *Context
	Verbose bool
}

type Context struct {
	Context    context.Context
	CancelFunc context.CancelFunc
	Channel    chan interface{}
}

func (cc *Context) Update(key interface{}, val interface{}) {
	cc.Context = context.WithValue(cc.Context, key, val)
}

func (cc *Context) Read(key interface{}) interface{} {
	return (cc.Context).Value(key)
}

func CreateContext() *Context {
	ctx, cancel := context.WithCancel(context.Background())

	return &Context{
		Context:    ctx,
		CancelFunc: cancel,
		Channel:    make(chan interface{}),
	}
}

func Run(params *Params) {
	var err error = params.Err
	var template string
	var ctx = params.Context

	if (ctx) == nil {
		panic("Cannot Run wihout Context")
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		template = fmt.Sprintf("%s:%d %%s\n", file, line)
	}

	if params.Err == nil && params.Handler != nil {
		err = params.Handler(ctx)
	}

	if err != nil {
		if params.Panic != nil {
			if params.Verbose {
				fmt.Printf("[VERBOSE] %s", fmt.Sprintf(template, "(Panic)"))
			}
			panic(params.Panic(err))
		}
		if params.Failed != nil {
			if params.Verbose {
				fmt.Printf("[VERBOSE] %s", fmt.Sprintf(template, "(Failed)"))
			}
			params.Failed(ctx)
		}
	} else {
		if params.Success != nil {
			if params.Verbose {
				fmt.Printf("[VERBOSE] %s", fmt.Sprintf(template, "(Success)"))
			}
			params.Success(ctx)
		}
	}
}
