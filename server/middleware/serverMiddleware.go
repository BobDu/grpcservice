package middleware

import (
	"context"
	pb "grpcservice"
)

type CacheServiceMiddleware struct {
	Next pb.CacheServiceServer
}

func BuildGetMiddleware(server pb.CacheServiceServer) pb.CacheServiceServer {
	throttleMiddleware := &ThrottleMiddleware{server}
	serviceMiddleware := &CacheServiceMiddleware{throttleMiddleware}
	return serviceMiddleware
}

func (serviceMiddleware *CacheServiceMiddleware) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	return serviceMiddleware.Next.Get(ctx, req)
}

func (serviceMiddleware *CacheServiceMiddleware) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	return serviceMiddleware.Next.Store(ctx, req)
}

func (serviceMiddleware *CacheServiceMiddleware) Dump(req *pb.DumpReq, server pb.CacheService_DumpServer) error {
	return serviceMiddleware.Next.Dump(req, server)
}
