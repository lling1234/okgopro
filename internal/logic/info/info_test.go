package info

import (
	"context"
	"okgopro/internal/model/entity"
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
		{name: "test1", args: args{ctx: context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sInfo{}
			s.GetXiaodao(tt.args.ctx)
		})
	}
}

func Test_sInfo_Getxk(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{ctx: context.Background()}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sInfo{}
			s.Getxk(tt.args.ctx)
		})
	}
}

func Test_sInfo_GetxkInfoDB(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{ctx: context.Background()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sInfo{}
			s.GetXiaodao(tt.args.ctx)
		})
	}
}

func TestQueryDBContent(t *testing.T) {
	type args struct {
		ctx context.Context
		in  entity.Content
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test1", args: args{ctx: context.Background(), in: entity.Content{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QueryDBContent(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("QueryDBContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
