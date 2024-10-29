// Package parser
// @Author: yinwei
// @File: parser_test.go
// @Version: 1.0.0
// @Date: 2024/3/13 17:25

package parser

import "testing"

func TestParser_Parse(t *testing.T) {
	type fields struct {
		PkgPath     string
		PkgName     string
		StructNames []string
		AllStructs  bool
		ExSuff      string
		WangSuff    string
	}
	type args struct {
		fname string
		isDir bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "excludGoFile", fields: fields{
			PkgPath:     "./testParse",
			PkgName:     "testParse",
			StructNames: nil,
			AllStructs:  true,
			ExSuff:      "",
			WangSuff:    ".pprof.go",
		}, args: args{
			fname: "./testParse",
			isDir: true,
		}, wantErr: false}, {name: "excludPprofGoFile", fields: fields{
			PkgPath:     "./testParse",
			PkgName:     "testParse",
			StructNames: nil,
			AllStructs:  true,
			ExSuff:      "",
			WangSuff:    ".go",
		}, args: args{
			fname: "./testParse",
			isDir: true,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				PkgPath:     tt.fields.PkgPath,
				PkgName:     tt.fields.PkgName,
				StructNames: tt.fields.StructNames,
				AllStructs:  tt.fields.AllStructs,
				ExSuff:      tt.fields.ExSuff,
				WantSuff:    tt.fields.WangSuff,
			}
			if err := p.Parse(tt.args.fname, tt.args.isDir); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
