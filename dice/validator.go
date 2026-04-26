package dice

import "fmt"

// Validate checks that count and sides are within valid ranges.
func Validate(count, sides int) error {
	if count < 1 || count > 100 {
		return fmt.Errorf("dice count must be between 1 and 100")
	}
	if sides < 1 || sides > 100 {
		return fmt.Errorf("dice sides must be between 1 and 100")
	}
	return nil
}
