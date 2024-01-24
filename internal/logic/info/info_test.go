package info

import (
	"context"
	"testing"
)

func Test_sInfo_GetXiaodao(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{ctx: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sInfo{}
			s.GetXiaodao(tt.args.ctx)
		})
	}
}
