package utils

import "testing"

func TestCreateSlug(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSlug(tt.args.input); got != tt.want {
				t.Errorf("CreateSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}
