package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	// ErrBadRouting - bad routing err
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
	// ErrInvalidArgument - invalid argument err
	ErrInvalidArgument = errors.New("invalid argument")
	// ErrInvalidArgument - invalid parameters err
	ErrInvalidParams = errors.New("invalid parameters")
	// ErrHTTPNotAllowed - HTTP Request Error
	ErrHTTPNotAllowed = errors.New("http not allowed")
	// ErrXFFProtoMissing - Specific HTTP Specification missing
	ErrXFFProtoMissing = errors.New("X-Forwarded-Proto header missing")
)

// MakeRoutes - creates routes for services
func MakeRoutes(r *mux.Router, s Service, logger log.Logger, middleware endpoint.Middleware) {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods(http.MethodGet).Path("/v1/physicians").Handler(kithttp.NewServer(
		middleware(MakePhysiciansExecuteEndpoint(s)),
		decodePhysicianExecutionRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/v1/schedule").Handler(kithttp.NewServer(
		middleware(MakeScheduleExecuteEndpoint(s)),
		decodeScheduleExecutionRequest,
		encodeResponse,
		options...,
	))
}

//decodePhysicianExecutionRequest - function to decode Job Request
func decodePhysicianExecutionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req PhysiciansRequest
	decoder := schema.NewDecoder()

	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		fmt.Println(err)
		return nil, ErrInvalidParams
	}

	return req, nil
}

//decodePhysicianExecutionRequest - function to decode Job Request
func decodeScheduleExecutionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ScheduleRequest
	decoder := schema.NewDecoder()

	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		fmt.Println(err)
		return nil, ErrInvalidParams
	}

	return req, nil
}

// encodeResponse - function to re-encode response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	switch err {
	case ErrInvalidArgument:
		code = http.StatusBadRequest
	case ErrInvalidParams:
		code = http.StatusBadRequest
	case ErrXFFProtoMissing:
		code = http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
