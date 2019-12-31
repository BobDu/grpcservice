package middleware

import (
	"context"
	"errors"
	pb "grpcservice"
	"log"
	"sync"
)

const (
	service_throttle = 5
)

var tm throttleMutex

type throttleMutex struct {
	mu sync.RWMutex
	throttle int
}

type ThrottleMiddleware struct {
	Next pb.CacheServiceServer
}

func (t *throttleMutex) getThrottle() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.throttle
}

func (t *throttleMutex) changeThrottle(delta int)  {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.throttle = t.throttle + delta
}

func (tg *ThrottleMiddleware) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	if tm.getThrottle() >= service_throttle {
		log.Printf("Get throttle=%v reached\n", tm.getThrottle())
		return nil, errors.New("Service throttle recached, please try later")
	} else {
		tm.changeThrottle(1)
		resp, err := tg.Next.Get(ctx, req)
		tm.changeThrottle(-1)
		return resp, err
	}
}

func (tg *ThrottleMiddleware) Store(ctx context.Context, req *pb.StoreReq) (*pb.StoreResp, error) {
	return tg.Next.Store(ctx, req)
}

func (tg *ThrottleMiddleware) Dump(req *pb.DumpReq, server pb.CacheService_DumpServer) error {
	return tg.Next.Dump(req, server)
}
