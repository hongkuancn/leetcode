package main

import (
	"reflect"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	type args struct {
		spells  []int
		potions []int
		success int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				spells:  []int{4, 1, 5},
				potions: []int{1, 4, 3, 2, 5},
				success: 7,
			},
			want: []int{4, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := successfulPairs(tt.args.spells, tt.args.potions, tt.args.success); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("successfulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
