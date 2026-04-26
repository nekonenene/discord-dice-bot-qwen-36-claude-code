package dice

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Roll generates count random numbers in range [1, sides] using crypto/rand.
func Roll(count, sides int) ([]int, error) {
	results := make([]int, count)
	max := big.NewInt(int64(sides))

	for i := 0; i < count; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, fmt.Errorf("failed to generate random number: %w", err)
		}
		results[i] = int(val.Int64()) + 1
	}

	return results, nil
}
