package geval

import (
	"log"
)

type Params struct {
	Err            error
	Success        func()
	Failed         func()
	Panic          func(v any) error
	Verbose        int
	SuccessMessage string
	FailedMessage  string
}

func Run(params *Params) {
	if params.Err != nil {
		if params.Verbose > 0 {
			log.Println(params.Err.Error())
		}
		if params.Panic != nil {
			panic(params.Panic(params.Err))
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
