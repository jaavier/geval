package geval

import (
	"log"
)

type Params struct {
	Err            error
	Success        func()
	Handler        func() error
	Failed         func()
	Panic          func(v any) error
	Verbose        int
	SuccessMessage string
	FailedMessage  string
}

func Run(params *Params) {
	var err error

	if params.Handler == nil {
		return
	}
	if params.Err == nil {
		err = params.Handler()
	}

	if err != nil {
		if params.Verbose > 0 {
			log.Println(err.Error())
		}
		if params.Panic != nil {
			panic(params.Panic(err))
		}
		if params.Failed != nil {
			params.Failed()
		}
	} else {
		if params.Success != nil {
			params.Success()
		}
	}
}
