package persona

import "context"

type CallRPC interface {
	Call(ctx context.Context) interface{}

	CallWithLocalCache(ctx context.Context) interface{}

	CallWithRedisCache(ctx context.Context) interface{}
}
