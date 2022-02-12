package stack

import (
	"github.com/tonoy30/goutils/common"
	"reflect"
	"testing"
)

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
		want []common.Item
	}{
		{
			name: "Case 1",
			want: []common.Item{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Demo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Demo() = %v, want %v", got, tt.want)
			}
		})
	}
}
