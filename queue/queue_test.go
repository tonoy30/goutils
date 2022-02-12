package queue

import (
	"github.com/tonoy30/goutils/common"
	"reflect"
	"testing"
)

func TestQueueDemo(t *testing.T) {
	tests := []struct {
		name string
		want []common.Item
	}{
		{
			name: "Case 1",
			want: []common.Item{2},
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
