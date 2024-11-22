// Code generated by Kitex v0.11.3. DO NOT EDIT.

package helloservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	hello "rpc-test/hello-server/kitex_gen/hello"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SayHello(ctx context.Context, req *hello.HelloRequest, callOptions ...callopt.Option) (r *hello.HelloResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kHelloServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kHelloServiceClient struct {
	*kClient
}

func (p *kHelloServiceClient) SayHello(ctx context.Context, req *hello.HelloRequest, callOptions ...callopt.Option) (r *hello.HelloResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SayHello(ctx, req)
}