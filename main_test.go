package main

import "testing"

func Test_triangular(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 3},
		{"3", args{3}, 6},
		{"4", args{4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangular(tt.args.n); got != tt.want {
				t.Errorf("triangular() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pentagonal(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"2", args{2}, 5},
		{"3", args{3}, 12},
		{"4", args{4}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pentagonal(tt.args.n); got != tt.want {
				t.Errorf("pentagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
