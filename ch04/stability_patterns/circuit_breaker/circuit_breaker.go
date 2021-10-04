package circuit_breaker

import "context"

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	return nil
}