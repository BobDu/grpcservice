/*
service包中的cacheClient 主要是用来调用服务端的函数用的
*/
package service

import (
	"context"
	pb "grpcservice"
)

/*
CacheClient是一个空结构， 为了实现CallGet()函数 即callGetter接口
*/
type CacheClient struct {
}

/*
主要的业务逻辑 修饰模式要完成的主要功能 其它功能都是为了对其修饰
*/
func (*CacheClient) CallGet(ctx context.Context, key string, csc pb.CacheServiceClient) ([]byte, error) {
	getReq := pb.GetReq{Key: key}
	getResp, err := csc.Get(ctx, &getReq)
	if err != nil {
		return nil, err
	}
	value := getResp.Value
	return value, err
}

func (*CacheClient) CallStore(key string, value []byte, client pb.CacheServiceClient) (*pb.StoreResp, error) {
	storeReq := pb.StoreReq{Key:key, Value:value}
	storeResp, err := client.Store(context.Background(), &storeReq)
	if err != nil {
		return nil, err
	}
	return storeResp, err
}
