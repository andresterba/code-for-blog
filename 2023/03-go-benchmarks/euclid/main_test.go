package main

import (
	"fmt"
	"testing"
)

func TestEuclid(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				a: 252,
				b: 105,
			},
			want: 21,
		},
		{
			name: "basic test",
			args: args{
				a: 105,
				b: 252,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Euclid(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Euclid() = %v, want %v", got, tt.want)
			}
		})
	}
}

var inputs = []struct {
	a int
	b int
}{
	{
		a: 252,
		b: 105,
	},
	{
		a: 1337,
		b: 105,
	},
}

func BenchmarkEuclid(b *testing.B) {
	for _, v := range inputs {
		b.Run(fmt.Sprintf("a:%d b:%d", v.a, v.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Euclid(v.a, v.b)
			}
		})
	}
}
