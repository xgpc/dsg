// package util
// @Author: dsg
// @Description:
// @File:  validate_test.go.go
// @Date: 2023/6/30 3:10

package util

import "testing"

func TestValidateIDCard(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"true",
			args{card: "530423199205270030"},
			true,
		}, {
			"false",
			args{card: "530423199205270031"},
			false,
		}, {
			"17位",
			args{card: "53042319920527003"},
			false,
		}, {
			"19位",
			args{card: "5304231992052700301"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIDCard(tt.args.card); got != tt.want {
				t.Errorf("ValidateIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
