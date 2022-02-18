// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package wallet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/venus/venus-shared/api"
)

const MajorVersion = 0
const APINamespace = "wallet.IFullAPI"
const MethodNamespace = "Filecoin"

// NewIFullAPIRPC creates a new httpparse jsonrpc remotecli.
func NewIFullAPIRPC(ctx context.Context, addr string, requestHeader http.Header, opts ...jsonrpc.Option) (IFullAPI, jsonrpc.ClientCloser, error) {
	endpoint, err := api.Endpoint(addr, MajorVersion)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid addr %s: %w", addr, err)
	}

	if requestHeader == nil {
		requestHeader = http.Header{}
	}
	requestHeader.Set(api.VenusAPINamespaceHeader, APINamespace)

	var res IFullAPIStruct
	closer, err := jsonrpc.NewMergeClient(ctx, endpoint, MethodNamespace, api.GetInternalStructs(&res), requestHeader, opts...)

	return &res, closer, err
}
