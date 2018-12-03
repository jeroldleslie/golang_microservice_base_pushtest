// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "fivekilometer/users/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	s := m.PathPrefix("/api/v1/users").Subrouter()
	makeHealthHandler(m, endpoints, options["Health"])
	makeCreateHandler(s, endpoints, options["Create"])
	makeGetByIdHandler(s, endpoints, options["GetById"])
	return m
}
