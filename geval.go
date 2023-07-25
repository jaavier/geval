package geval

import (
	"context"
	"log"
)

type Params struct {
	Err     error
	Success func(ctx context.Context)
	Failed  func(ctx context.Context)
	Handler func() (context.Context, error)
	Panic   func(v any) error
	Verbose int
	Context context.Context
}

func Run(params *Params) {
	var err error
	var ctx = context.TODO()

	if params.Handler == nil && params.Err == nil {
		return
	}

	if params.Err == nil {
		ctx, err = params.Handler()
	}

	if err != nil {
		if params.Verbose > 0 {
			log.Println(err.Error())
		}
		if params.Panic != nil {
			panic(params.Panic(err))
		}
		if params.Failed != nil {
			params.Failed(ctx)
		}
	} else {
		if params.Success != nil {
			params.Success(ctx)
		}
	}
}
