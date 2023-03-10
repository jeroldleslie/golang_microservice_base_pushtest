package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	"github.com/jeroldleslie/golang_microservice_base/users/pkg/io"
	"github.com/jeroldleslie/golang_microservice_base/users/pkg/service"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	User io.User `json:"user"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	T     io.User `json:"user"`
	Error error   `json:"error"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		t, error := s.Create(ctx, req.User)
		return CreateResponse{
			Error: error,
			T:     t,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Error
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, user io.User) (u io.User, error error) {
	request := CreateRequest{User: user}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).T, response.(CreateResponse).Error
}

// GetByIdRequest collects the request parameters for the GetById method.
type GetByIdRequest struct {
	Id string `json:"id"`
}

// GetByIdResponse collects the response parameters for the GetById method.
type GetByIdResponse struct {
	U     io.User `json:"user"`
	Error error   `json:"error"`
}

// MakeGetByIdEndpoint returns an endpoint that invokes GetById on the service.
func MakeGetByIdEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIdRequest)
		u, error := s.GetById(ctx, req.Id)
		return GetByIdResponse{
			Error: error,
			U:     u,
		}, nil
	}
}

// Failed implements Failer.
func (r GetByIdResponse) Failed() error {
	return r.Error
}

// GetById implements Service. Primarily useful in a client.
func (e Endpoints) GetById(ctx context.Context, id string) (u io.User, error error) {
	request := GetByIdRequest{Id: id}
	response, err := e.GetByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetByIdResponse).U, response.(GetByIdResponse).Error
}

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status := s.Health(ctx)
		return HealthResponse{Status: status}, nil
	}
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (status bool) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Status
}
