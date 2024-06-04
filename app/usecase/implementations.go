package usecase

import (
	"context"
	"fmt"
	"strings"

	singlefizzbuzz "github.com/kevinsudut/go-learn/app/single-fizz-buzz"
)

func (u usecase) UseCaseSingleFizzBuzzWithRange(ctx context.Context, from int64, to int64) (resp string, err error) {
	if from > to || to-from > 100 {
		return resp, fmt.Errorf("invalid range")
	}

	var (
		index   = int64(0)
		results = make([]string, to-from+1)
	)

	for i := from; i <= to; i++ {
		results[index] = singlefizzbuzz.SingleFizzBuzz(i)
		index++
	}

	return strings.Join(results, " "), nil
}
