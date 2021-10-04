package circuit_breaker

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var errIntentional = errors.New("intentional error")

func TestCircuitBreaker(t *testing.T) {

	t.Run("fail after 5 calls", func(t *testing.T) {
		callCount := 5

		circuit := failAfter(callCount)

		maxCalls := 6
		for currCallCount := 1; currCallCount <= maxCalls; currCallCount++ {
			got, err := circuit(context.Background())
			switch currCallCount {
			case maxCalls:
				// expect error
				require.Equal(t, errIntentional, err)
			default:
				require.Equal(t, "success", got)
			}
		}
	})

	t.Run("failed after 5 calls with circuit breaker", func(t *testing.T) {

	})
}

func failAfter(threshold int) Circuit {
	var count int
	// Using functional composition
	return func(ctx context.Context) (string, error) {
		count++
		if count > threshold {
			return "", errIntentional
		}
		return "success", nil
	}
}
