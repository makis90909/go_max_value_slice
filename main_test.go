package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateRandomElements(t *testing.T) {
	t.Run("Size < 0", func(t *testing.T) {
		result := generateRandomElements(-1)
		assert.Nil(t, result)
	})
	t.Run("Size = 0", func(t *testing.T) {
		result := generateRandomElements(0)
		assert.Nil(t, result)
	})
	t.Run("Size > 0", func(t *testing.T) {
		result := generateRandomElements(10)
		assert.Len(t, result, 10)
	})
}

func Test_maximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Nil", args{nil}, 0},
		{"Empty slice", args{[]int{}}, 0},
		{"1 element", args{[]int{7}}, 7},
		{"Negative", args{[]int{-8, -3, -11}}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
