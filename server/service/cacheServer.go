package service

import (
	"context"
	"fmt"
	pb "grpcservice"
)

type CacheService struct {
	Storage map[string][]byte
}

func (c *CacheService) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	fmt.Println("start server side Get called: ")
	key := req.GetKey()
	value := c.Storage[key]
	resp := &pb.GetResp{Value: value}
	fmt.Println("Get called with return of value: ", value)
	return resp, nil
}
