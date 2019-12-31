/*
服务重试
如果有错的话 就不断的去调用 rcg.Next.CallGet(ctx, key, csc)
直到返回正常或者达到重试上限  每次重试前都需要重试间隔
*/
package middleware

import (
	"context"
	pb "grpcservice"
	"log"
	"time"
)

const (
	retry_count = 3
	retry_interval = 200
)

type RetryCallGet struct {
	Next callGetter
}

func (rcg *RetryCallGet) CallGet(ctx context.Context, key string, client pb.CacheServiceClient) ([]byte, error) {
	var value []byte
	var err error
	for i := 0; i < retry_count; i++ {
		value, err = rcg.Next.CallGet(ctx, key, client)
		log.Printf("Retry number %v|error=%v", i+1, err)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(retry_interval) * time.Millisecond)
	}
	return value, err
}
