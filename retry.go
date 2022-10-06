package retry

import (
	"context"
	"time"
)

// Retry a function until it is successful, or the maximum number of attempts is reached. The function must correctly
// handle contexts which expire or are cancelled. Each attempt will receive a new context with a timeout of the
// duration built on the original parent.
func Retry(parent context.Context, duration time.Duration, attempts int, f func(ctx context.Context) error) (err error) {
	for i := 0; i < attempts; i++ {
		ctx, cancel := context.WithTimeout(parent, duration)

		if err = f(ctx); err == nil {
			cancel()
			return nil
		}

		cancel()
	}

	return
}

func RetryWithValue[T any](parent context.Context, duration time.Duration, attempts int, f func(ctx context.Context) (T, error)) (val T, err error) {
	for i := 0; i < attempts; i++ {
		ctx, cancel := context.WithTimeout(parent, duration)

		if val, err = f(ctx); err == nil {
			cancel()
			return val, nil
		}

		cancel()
	}

	return
}
