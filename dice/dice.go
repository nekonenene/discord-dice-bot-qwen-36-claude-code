// Package dice provides dice rolling functionality.
package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RollResult holds the outcome of a dice roll.
type RollResult struct {
	DiceCount int
	Sides     int
	Rolls     []int
	Total     int
}

// ParseAndRoll parses a dice expression like "4D6" and rolls the dice.
// The format is <count>D<sides>, e.g. "4D6" means 4 six-sided dice.
func ParseAndRoll(input string) (*RollResult, error) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "D")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid dice format: %q (expected NDSides)", input)
	}

	count, err := strconv.Atoi(parts[0])
	if err != nil || count < 1 || count > 100 {
		return nil, fmt.Errorf("dice count must be an integer between 1 and 100")
	}

	sides, err := strconv.Atoi(parts[1])
	if err != nil || sides < 1 || sides > 100 {
		return nil, fmt.Errorf("dice sides must be an integer between 1 and 100")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rolls := make([]int, count)
	total := 0
	for i := 0; i < count; i++ {
		roll := rng.Intn(sides) + 1
		rolls[i] = roll
		total += roll
	}

	return &RollResult{
		DiceCount: count,
		Sides:     sides,
		Rolls:     rolls,
		Total:     total,
	}, nil
}

// Format returns a human-readable string of the roll result.
func (r *RollResult) Format() string {
	parts := make([]string, len(r.Rolls))
	for i, v := range r.Rolls {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("%dD%d: [%s] 合計 %d",
		r.DiceCount, r.Sides, strings.Join(parts, ", "), r.Total)
}
