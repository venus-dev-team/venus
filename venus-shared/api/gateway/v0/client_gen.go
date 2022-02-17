// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package v0

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/venus/venus-shared/api"
)

// NewIGatewayRPC creates a new httpparse jsonrpc remotecli.
func NewIGatewayRPC(ctx context.Context, addr string, requestHeader http.Header, opts ...jsonrpc.Option) (IGateway, jsonrpc.ClientCloser, error) {
	if requestHeader == nil {
		requestHeader = http.Header{}
	}
	requestHeader.Set(api.VenusAPINamespaceHeader, "v0.IGateway")

	var res IGatewayStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Gateway", api.GetInternalStructs(&res), requestHeader, opts...)

	return &res, closer, err
}
