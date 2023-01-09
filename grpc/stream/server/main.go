/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/go-water/water"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	pb "stream.grpc/routeguide/proto"
)

var (
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

// GetFeature returns the feature at the given point.
func (h *handlers) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	_, response, err := h.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return response.(*pb.HelloReply), nil
}

type handlers struct {
	pb.UnimplementedGreeterServer
	sayHello water.Handler
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, NewServer())
	grpcServer.Serve(lis)
}

func NewServer() pb.GreeterServer {
	return &handlers{
		sayHello: water.NewHandler(HelloService{}),
	}
}

type HelloService struct {
	serverBeforeAfter
}

func (svc HelloService) Handle(ctx context.Context, req *pb.HelloRequest) (interface{}, error) {
	grpclog.Infof("%+v", req)
	return nil, nil
}

func (svc HelloService) Endpoint() water.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HelloRequest)
		return svc.Handle(ctx, req)
	}
}

func (HelloService) DecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	return req.(*pb.HelloRequest), nil
}

func (HelloService) EncodeResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	reply := new(pb.HelloReply)
	reply.Message = "min"
	return reply, nil
}

func (svc HelloService) GetLogger() grpclog.DepthLoggerV2 {
	return grpclog.Component(svc.Name())
}

func (svc HelloService) Name() string {
	return svc.name(svc)
}

type serverBeforeAfter struct{}

func (s serverBeforeAfter) ServerBefore() water.ServerOption {
	return water.ServerBefore(s.serverBefore)
}

func (s serverBeforeAfter) ServerAfter() water.ServerOption {
	return water.ServerAfter(s.serverAfter)
}

func (s serverBeforeAfter) name(srv interface{}) string {
	fullName := fmt.Sprintf("%T", srv)
	index := strings.LastIndex(fullName, ".")
	name := fullName[index+1:]

	return name
}

func (serverBeforeAfter) serverBefore(ctx context.Context, md metadata.MD) context.Context {
	return ctx
}

func (serverBeforeAfter) serverAfter(ctx context.Context, header *metadata.MD, trailer *metadata.MD) context.Context {
	return ctx
}
