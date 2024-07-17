package dis_doi

import (
	"context"
	"testing"
)

func TestIsReserved(t *testing.T) {
	initWhiteList("whitelist")
	type args struct {
		ctx  context.Context
		doi  string
		zone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"reserved",
			args{
				context.Background(),
				"test.viv.cn.",
				"viv.cn",
			},
			true,
		},
		{
			"not-re",
			args{
				context.Background(),
				"testfill.viv.cn.",
				"viv.cn",
			},
			false,
		},
		{
			"not-re",
			args{
				context.Background(),
				"zhiku.viv.cn.",
				"viv.cn",
			},
			false,
		},
		{
			"not-re",
			args{
				context.Background(),
				"fuxi.viv.cn.",
				"viv.cn",
			},
			false,
		},
		{
			"re-1",
			args{
				context.Background(),
				"fuxi1.viv.cn.",
				"viv.cn",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReserved(tt.args.ctx, tt.args.doi, tt.args.zone); got != tt.want {
				t.Errorf("IsReserved() = %v, want %v", got, tt.want)
			}
		})
	}
}
