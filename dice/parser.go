package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

var diceRegex = regexp.MustCompile(`(?i)^(\d+)D(\d+)$`)

// Parse parses dice notation (e.g., "4D6", "2D100") and returns count and sides.
func Parse(input string) (count, sides int, err error) {
	matches := diceRegex.FindStringSubmatch(input)
	if matches == nil {
		return 0, 0, fmt.Errorf("invalid dice notation: %q", input)
	}

	count, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid dice count: %w", err)
	}

	sides, err = strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid dice sides: %w", err)
	}

	return count, sides, nil
}
