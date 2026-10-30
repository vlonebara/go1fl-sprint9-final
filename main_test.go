package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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
				assert.Nil(t, got)
				return
			}

			assert.Len(t, got, tt.size)
			for i, n := range got {
				assert.GreaterOrEqual(t, n, 0, "element %d is negative", i)
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
			name: "single element",
			data: []int{42},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaximumEmptyInput(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{name: "nil slice", data: nil},
		{name: "empty slice", data: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Zero(t, maximum(tt.data))
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
		{
			name: "fewer elements than chunks",
			data: []int{5, 9, 7},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMaxChunksEmptyInput(t *testing.T) {
	assert.Zero(t, maxChunks(nil))
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	sizes := []int{1, 3, 7, 8, 9, 16, 31, 100, 1001}

	for _, size := range sizes {
		t.Run("size_"+strconv.Itoa(size), func(t *testing.T) {
			data := make([]int, size)

			for i := range data {
				data[i] = rng.Intn(2001) - 1000
			}

			want := maximum(data)
			got := maxChunks(data)

			assert.Equal(t, want, got, "size %d", size)
		})
	}
}
