package middleware

import (
	"context"
	pb "grpcservice"
	"log"
	"time"
)

const (
	get_timeout = 200
)

type TimeoutCallGet struct {
	Next callGetter
}

func (tcg *TimeoutCallGet) CallGet(ctx context.Context, key string, client pb.CacheServiceClient) ([]byte, error) {
	var cancelFunc context.CancelFunc
	var ch = make(chan bool)
	var value []byte
	var err error
	ctx, cancelFunc = context.WithTimeout(ctx, get_timeout*time.Millisecond)
	go func() {
		value, err = tcg.Next.CallGet(ctx, key, client)
		ch <- true
	}()
	select {
		case <- ctx.Done():
			log.Fatalln("ctx timeout")
			cancelFunc()
			err = ctx.Err()
		case <- ch:
			log.Println("call finished normally")
	}
	return value, err
}
