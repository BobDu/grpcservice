package middleware

import (
	"context"
	pb "grpcservice"
)

/*
callGetter 接口 修饰模式的核心 每个修饰功能都需要实现这个接口
接口中只要求提供一个函数CallGet 这个函数会被每个修饰功能不断调用 该函数的签名是按照业务函数来定义
*/
type callGetter interface {
	CallGet(ctx context.Context, key string, client pb.CacheServiceClient) ([]byte, error)
}

/*
CallGetMiddleware 结构struct
里面只有一个成员Next 类型为修饰模式接口 callGetter 这样就可以通过Next来调用下一个修饰功能
每一个修饰功能结构都会有一个相应的修饰结构  需要把他们串起来 才能完成依次叠加调用
*/
type CallGetMiddleware struct {
	Next callGetter
}

func BuildGetMiddleware(cc callGetter) callGetter {
	//cbcg := CircuitBreakerCallGet{cc}
	tcg := &TimeoutCallGet{cc}
	rcg := &RetryCallGet{tcg}
	return rcg
}

func (cg *CallGetMiddleware) CallGet(ctx context.Context, key string, client pb.CacheServiceClient) ([]byte, error) {
	return cg.Next.CallGet(ctx, key, client)
}