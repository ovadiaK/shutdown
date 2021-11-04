package input

import (
	"context"
	"fmt"
	"time"
)

const limit = 10000000

type Input interface {
	GetInts(ctx context.Context) []int
}

type fake struct {
	n int
}

func New() *fake {
	return &fake{n: 1}
}

func (f *fake) GetInts(ctx context.Context) []int {
	result := make([]int, 0, limit)
	for i := 0; i < limit; i++ {
		select {
		case <-time.After(time.Nanosecond):
			result = append(result, 1)
		case <-ctx.Done():
			fmt.Printf("current input: %d\n", f.n)
			return result
		}
		f.n += 1
	}
	return result
}
