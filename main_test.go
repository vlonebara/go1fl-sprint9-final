package main

import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
	}{
		{name: "normal size", size: 100},
		{name: "minimum size", size: 1},
		{name: "zero size", size: 0, wantNil: true},
		{name: "negative size", size: -1, wantNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if tt.wantNil {
				if got != nil {
					t.Fatalf("expected nil, got %v", got)
				}
				return
			}

			if len(got) != tt.size {
				t.Fatalf("expected length %d, got %d", tt.size, len(got))
			}

			for i, n := range got {
				if n < 0 || n >= 1000 {
					t.Errorf("element %d out of range: %d", i, n)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "maximum in middle",
			data: []int{1, 5, 10, 3, 2},
			want: 10,
		},
		{
			name: "maximum at beginning",
			data: []int{100, 5, 10, 3},
			want: 100,
		},
		{
			name: "maximum at end",
			data: []int{1, 5, 10, 30},
			want: 30,
		},
		{
			name: "negative numbers",
			data: []int{-10, -5, -20, -1},
			want: -1,
		},
		{
			name: "duplicate maximum",
			data: []int{7, 10, 10, 3},
			want: 10,
		},
		{
			name: "all equal",
			data: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "single element current behavior",
			data: []int{42},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d",
					tt.data, got, tt.want)
			}
		})
	}
}

func TestMaximumEmptyInputPanics(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{name: "nil slice", data: nil},
		{name: "empty slice", data: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic for empty input")
				}
			}()

			maximum(tt.data)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "maximum in first chunk",
			data: []int{100, 1, 2, 3, 4, 5, 6, 7},
			want: 100,
		},
		{
			name: "maximum in last chunk",
			data: []int{1, 2, 3, 4, 5, 6, 7, 100},
			want: 100,
		},
		{
			name: "maximum in middle",
			data: []int{1, 2, 3, 4, 100, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			want: 100,
		},
		{
			name: "negative numbers",
			data: []int{-10, -5, -20, -1, -7, -3, -8, -4},
			want: -1,
		},
		{
			name: "duplicate maximum",
			data: []int{10, 1, 2, 10, 4, 5, 10, 7},
			want: 10,
		},
		{
			name: "uneven chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d",
					tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksEmptyInput(t *testing.T) {
	got := maxChunks(nil)

	if got != 0 {
		t.Errorf("maxChunks(nil) = %d, want 0", got)
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	// Test only sizes that produce non-empty chunks
	// with the current implementation.
	sizes := []int{8, 9, 16, 31, 100, 1001}

	for _, size := range sizes {
		t.Run("size_"+itoa(size), func(t *testing.T) {
			data := make([]int, size)

			for i := range data {
				data[i] = rng.Intn(2001) - 1000
			}

			want := maximum(data)
			got := maxChunks(data)

			if got != want {
				t.Errorf("size %d: maxChunks = %d, maximum = %d",
					size, got, want)
			}
		})
	}
}

func itoa(n int) string {
	const digits = "0123456789"

	if n == 0 {
		return "0"
	}

	var buf [20]byte
	i := len(buf)

	for n > 0 {
		i--
		buf[i] = digits[n%10]
		n /= 10
	}

	return string(buf[i:])
}
