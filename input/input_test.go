package input

import (
	"fmt"
	"os"
	"testing"
)

func Test_a1(t *testing.T) {
	a := os.Getenv("asergasdg")
	fmt.Println(a)
}

func Test_withUnitToInt(t *testing.T) {
	type args struct {
		withUnit string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GB",
			args: args{
				withUnit: "10GB",
			},
			want: 10000000000,
		},
		{
			name: "KiB",
			args: args{
				withUnit: "10KiB",
			},
			want: 10240,
		},
		{
			name: "MiB",
			args: args{
				withUnit: "1MiB",
			},
			want: 1048576,
		},
		{
			name: "b",
			args: args{
				withUnit: "1b",
			},
			want: 1,
		},
		{
			name: "kib_space",
			args: args{
				withUnit: "3 kib",
			},
			want: 3072,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withUnitToInt(tt.args.withUnit); got != tt.want {
				t.Errorf("withUnitToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
