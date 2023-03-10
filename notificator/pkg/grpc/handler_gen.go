// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"

	"github.com/jeroldleslie/golang_microservice_base/notificator/pkg/endpoint"
	"github.com/jeroldleslie/golang_microservice_base/notificator/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	notify grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.NotificatorServer {
	return &grpcServer{
		notify: makeNotifyHandler(endpoints, options["Notify"]),
	}
}
