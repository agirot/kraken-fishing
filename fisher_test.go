package main

import (
	"testing"
)

func Test_buy(t *testing.T) {
	maxHoldCount = 2
	type args struct {
		target             float64
		currentClosedPrice float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "wait to lower",
			args: args{
				target:             12,
				currentClosedPrice: 15,
			},
			want: false,
		},
		{
			name: "wait to lower",
			args: args{
				target:             12,
				currentClosedPrice: 14,
			},
			want: false,
		},
		{
			name: "drop, don't buy",
			args: args{
				target:             12,
				currentClosedPrice: 8,
			},
			want: false,
		},
		{
			name: "wait to lower again",
			args: args{
				target:             12,
				currentClosedPrice: 7,
			},
			want: false,
		},
		{
			name: "wait to lower again",
			args: args{
				target:             12,
				currentClosedPrice: 6,
			},
			want: false,
		},
		{
			name: "hold",
			args: args{
				target:             12,
				currentClosedPrice: 7,
			},
			want: false,
		},
		{
			name: "hold",
			args: args{
				target:             12,
				currentClosedPrice: 8,
			},
			want: false,
		},
		{
			name: "buy",
			args: args{
				target:             12,
				currentClosedPrice: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buy(tt.args.target, tt.args.currentClosedPrice); got != tt.want {
				t.Errorf("fisher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sell(t *testing.T) {
	maxHoldCount = 2
	type args struct {
		target             float64
		currentClosedPrice float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "wait to higher",
			args: args{
				target:             12,
				currentClosedPrice: 11,
			},
			want: false,
		},
		{
			name: "wait to higher",
			args: args{
				target:             12,
				currentClosedPrice: 13,
			},
			want: false,
		},
		{
			name: "drop, don't sell",
			args: args{
				target:             12,
				currentClosedPrice: 10,
			},
			want: false,
		},
		{
			name: "wait to higher again",
			args: args{
				target:             12,
				currentClosedPrice: 16,
			},
			want: false,
		},
		{
			name: "wait to higher again",
			args: args{
				target:             12,
				currentClosedPrice: 17,
			},
			want: false,
		},
		{
			name: "hold",
			args: args{
				target:             12,
				currentClosedPrice: 15,
			},
			want: false,
		},
		{
			name: "hold",
			args: args{
				target:             12,
				currentClosedPrice: 14,
			},
			want: false,
		},
		{
			name: "sell",
			args: args{
				target:             12,
				currentClosedPrice: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sell(tt.args.target, tt.args.currentClosedPrice); got != tt.want {
				t.Errorf("fisher() = %v, want %v", got, tt.want)
			}
		})
	}
}
