// Package gen
// @Author: yinwei
// @File: encode_test.go
// @Version: 1.0.0
// @Date: 2024/12/3 18:10

package gen

import (
	"testing"
)

func Test_generator_parseCrossTag(t *testing.T) {

	type args struct {
		fieldName string
		tag       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "jiexi", args: args{
			fieldName: "CrossVad",
			tag:       "A=int_lt:12;int_gt:12|C=float_lt:0.3|CrossVad=required",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &generator{}
			got, err := g.parseCrossTag(tt.args.fieldName, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCrossTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}
