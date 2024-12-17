package utils

import "testing"

func TestStringToUint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUint(tt.args.s); got != tt.want {
				t.Errorf("StringToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
