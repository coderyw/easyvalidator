/**
 * @Author: yinwei
 * @Description:
 * @File: generator_test.go
 * @Version: 1.0.0
 * @Date: 2023/4/24 11:05
 */
package gen

import (
	"github.com/coderyw/easyvalidator/gen/test/model"
	"os"
	"testing"
)

type EasyMAP_exporter_TestStruct *model.TestStruct
type EasyMAP_exporter_Fs *model.Fs

func Test_generator_Run(t *testing.T) {
	type fields struct {
		generator string
		pkg       string
		obj       interface{}
		outFile   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "TestStruct", fields: fields{
			generator: "TestStruct.go",
			pkg:       "github.com/coderyw/easyvalidator/test/model",
			obj:       EasyMAP_exporter_TestStruct(nil),
			outFile:   "test/model/TestStruct_easyvalidator.go",
		}, wantErr: false},
		{name: "TestFs", fields: fields{
			generator: "Fs.go",
			pkg:       "github.com/coderyw/easyvalidator/test/model",
			obj:       EasyMAP_exporter_Fs(nil),
			outFile:   "test/model/Fs_easyvalidator.go",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator(tt.fields.generator)
			g.SetPkg("model", tt.fields.pkg)
			g.Add(tt.fields.obj)
			f, err := os.Create(tt.fields.outFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			{
				defer f.Close()
				g.Run(f)
			}
		})
	}
}

func TestUnMarshalMapInterface(t *testing.T) {
	//a := model.OutModel{}
	//a.UnEasyMap = model1.UnEasyMap{
	//	A: 1,
	//}
	//a.PtrMap = &model1.UnEasyMap{A: 2}
	//m, _ := a.MarshalMap()
	//fmt.Println(m)
	//b := &model.OutModel{}
	//err := b.UnMarshalMapInterface(m)
	//fmt.Println(err)
	//fmt.Println(b)
}
