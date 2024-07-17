package dis_doi

import (
	"context"
	"testing"
)

func Test_initSensitiveWords(t *testing.T) {
	type args struct {
		sourceName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"normal",
			args{
				"sensitive_words_test",
			},
			false,
		},
		{
			"not exist",
			args{
				"sensitive_words_test_ne",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if num, err := initSensitiveWords(tt.args.sourceName); (err != nil) != tt.wantErr {
				t.Errorf("initSensitiveWords() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				t.Logf("load num:%d", num)
			}
		})
	}
}

func TestIsValidated(t *testing.T) {
	n, _ := initSensitiveWords("sensitive_words_test")
	t.Logf("sensitive_words_test:%d", n)
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
			"hit",
			args{
				context.Background(),
				"xitele.viv.cn",
				"viv.cn",
			},
			false,
		},
		{
			"hit-2",
			args{
				context.Background(),
				"xitele.viv.cn.",
				"viv.cn",
			},
			false,
		},
		{
			"hit-3",
			args{
				context.Background(),
				"xitele.viv.cn.",
				"viv.cn.",
			},
			false,
		},
		{
			"hit-4",
			args{
				context.Background(),
				"希特勒.viv.cn.",
				"viv.cn.",
			},
			false,
		},
		{
			"not-hit",
			args{
				context.Background(),
				"希特勒1.viv.cn.",
				"viv.cn.",
			},
			true,
		},
		{
			"should-not-hit",
			args{
				context.Background(),
				"kxz34m03uc.viv.cn.",
				"viv.cn.",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidated(tt.args.ctx, tt.args.doi, tt.args.zone); got != tt.want {
				t.Errorf("IsValidated() = %v, want %v", got, tt.want)
			}
		})
	}
}
