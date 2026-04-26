package dice

import (
	"strings"
	"testing"
)

func TestParseAndRoll(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantCount int
		wantSides int
		wantErr   bool
	}{
		{
			name:      "4D6",
			input:     "4D6",
			wantCount: 4,
			wantSides: 6,
			wantErr:   false,
		},
		{
			name:      "2D100",
			input:     "2D100",
			wantCount: 2,
			wantSides: 100,
			wantErr:   false,
		},
		{
			name:      "1D1",
			input:     "1D1",
			wantCount: 1,
			wantSides: 1,
			wantErr:   false,
		},
		{
			name:      "100D100",
			input:     "100D100",
			wantCount: 100,
			wantSides: 100,
			wantErr:   false,
		},
		{
			name:    "0D6",
			input:   "0D6",
			wantErr: true,
		},
		{
			name:    "101D6",
			input:   "101D6",
			wantErr: true,
		},
		{
			name:    "4D0",
			input:   "4D0",
			wantErr: true,
		},
		{
			name:    "4D101",
			input:   "4D101",
			wantErr: true,
		},
		{
			name:    "invalid",
			input:   "abc",
			wantErr: true,
		},
		{
			name:    "no D separator",
			input:   "46",
			wantErr: true,
		},
		{
			name:    "negative count",
			input:   "-1D6",
			wantErr: true,
		},
		{
			name:    "negative sides",
			input:   "4D-6",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseAndRoll(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if result.DiceCount != tt.wantCount {
				t.Errorf("DiceCount = %d, want %d", result.DiceCount, tt.wantCount)
			}
			if result.Sides != tt.wantSides {
				t.Errorf("Sides = %d, want %d", result.Sides, tt.wantSides)
			}
			if len(result.Rolls) != tt.wantCount {
				t.Errorf("Roll length = %d, want %d", len(result.Rolls), tt.wantCount)
			}
			for _, r := range result.Rolls {
				if r < 1 || r > tt.wantSides {
					t.Errorf("roll value %d out of range [1, %d]", r, tt.wantSides)
				}
			}
			if result.Total != 0 {
				for _, r := range result.Rolls {
					result.Total += r
				}
			}
		})
	}
}

func TestFormat(t *testing.T) {
	result := &RollResult{
		DiceCount: 4,
		Sides:     6,
		Rolls:     []int{3, 5, 1, 6},
		Total:     15,
	}
	got := result.Format()
	if !strings.Contains(got, "4D6") {
		t.Errorf("Format() = %q, want to contain '4D6'", got)
	}
	if !strings.Contains(got, "合計 15") {
		t.Errorf("Format() = %q, want to contain '合計 15'", got)
	}
}
