package dice

import (
	"fmt"
	"strconv"
	"strings"
)

// Result holds the outcome of a dice roll.
type Result struct {
	Count  int
	Sides  int
	Values []int
	Total  int
}

// String returns the formatted result: "[v1, v2, ..., vn] 合計: total"
func (r Result) String() string {
	parts := make([]string, len(r.Values))
	for i, v := range r.Values {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("[%s] 合計: %d", strings.Join(parts, ", "), r.Total)
}

// RollNotation parses, validates, rolls, and formats dice notation.
func RollNotation(input string) (Result, error) {
	count, sides, err := Parse(input)
	if err != nil {
		return Result{}, err
	}

	if err := Validate(count, sides); err != nil {
		return Result{}, err
	}

	values, err := Roll(count, sides)
	if err != nil {
		return Result{}, err
	}

	total := 0
	for _, v := range values {
		total += v
	}

	return Result{
		Count:  count,
		Sides:  sides,
		Values: values,
		Total:  total,
	}, nil
}
