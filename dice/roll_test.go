package dice

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input     string
		wantCount int
		wantSides int
		wantErr   bool
	}{
		{"4D6", 4, 6, false},
		{"2D100", 2, 100, false},
		{"1D20", 1, 20, false},
		{"100D6", 100, 6, false},
		{"1d6", 1, 6, false},
		{"d6", 0, 0, true},
		{"4d", 0, 0, true},
		{"4D", 0, 0, true},
		{"", 0, 0, true},
		{"abc", 0, 0, true},
		{"4D6s", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			count, sides, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if count != tt.wantCount || sides != tt.wantSides {
					t.Errorf("Parse(%q) = %d, %d, want %d, %d", tt.input, count, sides, tt.wantCount, tt.wantSides)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		count   int
		sides   int
		wantErr bool
	}{
		{1, 1, false},
		{100, 100, false},
		{1, 100, false},
		{100, 1, false},
		{0, 6, true},
		{101, 6, true},
		{1, 0, true},
		{1, 101, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			err := Validate(tt.count, tt.sides)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate(%d, %d) error = %v, wantErr %v", tt.count, tt.sides, err, tt.wantErr)
			}
		})
	}
}

func TestRoll_Bounds(t *testing.T) {
	values, err := Roll(1000, 6)
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range values {
		if v < 1 || v > 6 {
			t.Errorf("value[%d] = %d, want [1,6]", i, v)
		}
	}
}

func TestRollNotation(t *testing.T) {
	r, err := RollNotation("3D6")
	if err != nil {
		t.Fatal(err)
	}
	if r.Count != 3 || r.Sides != 6 {
		t.Errorf("got %dD%d, want 3D6", r.Count, r.Sides)
	}
	if r.Total < 3 || r.Total > 18 {
		t.Errorf("total %d out of range [3,18]", r.Total)
	}
}

func TestResult_String(t *testing.T) {
	r := Result{Count: 4, Sides: 6, Values: []int{3, 5, 1, 6}, Total: 15}
	got := r.String()
	want := "[3, 5, 1, 6] 合計: 15"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
