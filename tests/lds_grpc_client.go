// Copyright 2018 Google Cloud Platform Proxy Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8790, "LDS port")
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	addr := fmt.Sprintf("localhost:%d", *port)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		glog.Exitf("failed to connect to server: %v", err)
	}

	client := v2.NewListenerDiscoveryServiceClient(conn)
	ctx := context.Background()

	req := &v2.DiscoveryRequest{
		Node: &core.Node{
			Id: "api_proxy",
		},
	}
	resp := &v2.DiscoveryResponse{}
	if resp, err = client.FetchListeners(ctx, req); err != nil {
		glog.Exitf("discovery: %v", err)
	}

	fmt.Println(resp)
	// All fine.
	return
}
