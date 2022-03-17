package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Endpoints struct {
	GetPhsyiciansExecuteEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service, logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetPhsyiciansExecuteEndpoint: wrapEndpoint(MakePhysiciansExecuteEndpoint(svc), middlewares),
	}
}

// MakePhysiciansExecuteEndpoint - function that returns the endpoint for getting physicians
func MakePhysiciansExecuteEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PhysiciansRequest)
		if !ok {
			return nil, ErrBadRouting
		}
		res, err := svc.PhysiciansExecute(ctx, req)
		return res, err
	}
}

// MakeScheduleExecuteEndpoint - function that returns getting the schedule of a physician
func MakeScheduleExecuteEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ScheduleRequest)
		if !ok {
			return nil, ErrBadRouting
		}
		res, err := svc.ScheduleExecute(ctx, req)
		return res, err
	}
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
	  e = m(e)
	}
	
	return e
}
