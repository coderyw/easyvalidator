/**
 * @Author: yinwei
 * @Description:
 * @File: encode
 * @Version: 1.0.0
 * @Date: 2022/10/29 14:31
 */
package gen

import (
	"bytes"
	"fmt"
	vd "github.com/coderyw/easyvalidator/validator"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type validator struct {
	key   string
	value string
}

func (g *generator) encode(t reflect.Type) error {
	if t.Kind() == reflect.Struct {
		err := g.encodeStruct(t)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("暂时只支持struct")
	}
}

func (g *generator) encodeStruct(t reflect.Type) error {
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out)

	buffer := &bytes.Buffer{}

	fmt.Fprintln(buffer, fmt.Sprintf("func (this *%v) Validate() error {", t.Name()))
	var (
		field reflect.StructField
		fv    reflect.Type
	)
	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		fv = field.Type
		data, err := g.encodeField(t.Name(), fv, field, false)
		if err != nil {
			return err
		}
		if len(data) != 0 {
			buffer.Write(data)
		}
	}
	fmt.Fprintln(g.out, buffer.String())
	fmt.Fprintln(g.out, fmt.Sprintf("\treturn  nil"))
	fmt.Fprintln(g.out, fmt.Sprintf("}"))
	return nil
}

func (g *generator) getTag(s reflect.StructField) string {
	tag := s.Tag.Get(tag)
	if tag == "" {
		tag = s.Tag.Get(jsTag)
		arr := strings.Split(tag, ",")
		tag = arr[0]
	}
	if tag == "" {
		tag = s.Name
	}
	return tag
}

func (g *generator) encodeField(structName string, fv reflect.Type, field reflect.StructField, isPtr bool) ([]byte, error) {
	vad := field.Tag.Get(TAG_VALIDATOR)
	if len(vad) == 0 {
		return nil, nil
	}

	valids, err := g.parseTag(vad)
	if err != nil {
		return nil, err
	}

	bf := new(bytes.Buffer)

	// 遍历tag，生成对应的验证规则
	for _, v := range valids {
		data, err := g.make(structName, fv, field, v)
		if err != nil {
			return nil, err
		}
		bf.Write(data)
	}
	return bf.Bytes(), nil
}

func (g *generator) make(structName string, fv reflect.Type, field reflect.StructField, valid validator) ([]byte, error) {
	name := field.Name
	switch valid.key {
	case vd.VALIDATOR_REGEX:
		return g.regexValid(structName, name, valid.key, valid.value)
	case vd.VALIDATOR_INT_GT, vd.VALIDATOR_INT_GTE, vd.VALIDATOR_INT_LT, vd.VALIDATOR_INT_LTE:
		return g.intValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_MSG_EXISTS:
		return g.msgExistsValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_REQUIRED:
		return g.requiredValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_FLOAT_GT, vd.VALIDATOR_FLOAT_LT, vd.VALIDATOR_FLOAT_GTE, vd.VALIDATOR_FLOAT_LTE:
		return g.floatValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_LEN_GT, vd.VALIDATOR_LEN_LT, vd.VALIDATOR_LEN_EQ:
		return g.lenValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_UUID_VER:
		return g.uuidVer(fv, field, valid.key, valid.value)
	default:
		return nil, fmt.Errorf("错误的验证类型")
	}
}

func (g *generator) uuidVer(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {
	if fv.Kind() != reflect.String {
		return nil, fmt.Errorf("uuid var only support string！")
	}
	var (
		val int64
		err error
	)
	if value != "" {
		val, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	switch val {
	case 0, 1, 2, 3, 4, 5:
	default:
		return nil, fmt.Errorf("uuid_ver’s version only support 0~5.")
	}
	valib := fmt.Sprintf("_regex_uuidver%v", val)
	if _, ok := g.topVar[valib]; !ok {
		if val == 0 {
			g.topVar[valib] = `regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")`
		} else {
			g.topVar[valib] = fmt.Sprintf(`regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[%v][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")`, val)
		}
	}
	bf := new(bytes.Buffer)
	b := `	if %v.MatchString(this.%v) {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute %v must conform to the format of UUID version %v. "))
	}`
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	fmt.Fprintln(bf, fmt.Sprintf(b, valib, field.Name, field.Name, field.Name, val))
	return bf.Bytes(), nil
}

func (g *generator) lenValid(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {
	switch fv.Kind() {
	case reflect.Map, reflect.String, reflect.Array, reflect.Slice:

	default:
		return nil, fmt.Errorf("len validator only support map,string,array,slice")
	}
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}

	var (
		op, notice string
	)

	switch key {
	case vd.VALIDATOR_LEN_GT:
		op = "<="
		notice = "greater than"
	case vd.VALIDATOR_LEN_LT:
		op = ">="
		notice = "less than"
	case vd.VALIDATOR_LEN_EQ:
		op = "!="
		notice = "equal to"
	}

	bf := new(bytes.Buffer)
	b := `	if len(this.%v) %v %v{
		return validator_helper.FieldError("%v", fmt.Errorf("The length of the value of attribute %v must be %v %v. "))
	}`
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	fmt.Fprintln(bf, fmt.Sprintf(b, field.Name, op, val, field.Name, field.Name, notice, value))
	return bf.Bytes(), nil
}

func (g *generator) floatValid(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {
	switch fv.Kind() {
	case reflect.Float64, reflect.Float32:
	default:
		return nil, fmt.Errorf("only support type float")
	}

	fieldName := field.Name

	vi, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	var (
		op     string
		notice string
	)
	switch key {
	case vd.VALIDATOR_FLOAT_GT:
		op = "<="
		notice = "greater than"
	case vd.VALIDATOR_FLOAT_GTE:
		op = "<"
		notice = "greater than or equal to"
	case vd.VALIDATOR_FLOAT_LT:
		op = ">="
		notice = "less than"
	case vd.VALIDATOR_FLOAT_LTE:
		op = ">"
		notice = "less than or equal to"
	}
	b := `	if !(this.%v %v %v) {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must be %v '%v'. "))
	}
`
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	fmt.Fprintln(bf, fmt.Sprintf(b, fieldName, op, vi, fieldName, fieldName, notice, vi))
	return bf.Bytes(), nil
}

func (g *generator) requiredValid(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {
	if value == "false" || value == "0" {
		return nil, nil
	}
	var b string
	switch fv.Kind() {
	case reflect.Ptr:
		b = `	if nil == this.%v{
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must required. "))
	}`
		b = fmt.Sprintf(b, field.Name, field.Name, field.Name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64, reflect.Float32:
		b = `	if 0 == this.%v{
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must required. "))
	}`
		b = fmt.Sprintf(b, field.Name, field.Name, field.Name)
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		b = `	if len(this.%v) ==0{
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must required. "))
	}`
		b = fmt.Sprintf(b, field.Name, field.Name, field.Name)
	default:
		return nil, fmt.Errorf("required only support int,int8,int16,int32,int64,uin,uin8,uin16,uin32,uin64,float32,float64,array,map,pointer,slice,string")
	}
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	fmt.Fprintln(bf, b)
	return bf.Bytes(), nil
}

func (g *generator) msgExistsValid(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {

	if fv.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("msg exist 只对指针类型生效")
	}

	fieldName := field.Name
	bf := new(bytes.Buffer)
	b := `	if nil == this.%v {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must exist. "))
	}`
	b2 := `	if this.%v != nil{
		if err:= validator_helper.CallValidatorIfExists(this.%v); err!= nil{
			return validator_helper.FieldError("%v", err)
		}
	}`
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	fmt.Fprintln(bf, fmt.Sprintf(b, fieldName, fieldName, fieldName))
	fmt.Fprintln(bf, fmt.Sprintf(b2, fieldName, fieldName, fieldName))
	return bf.Bytes(), nil
}

// int
func (g *generator) intValid(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, error) {
	switch fv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		return nil, fmt.Errorf("only support type int")

	}
	fieldName := field.Name

	vi, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, err
	}

	var (
		op     string
		notice string
	)
	switch key {
	case vd.VALIDATOR_INT_GT:
		op = "<="
		notice = "greater than"
	case vd.VALIDATOR_INT_GTE:
		op = "<"
		notice = "greater than or equal to"
	case vd.VALIDATOR_INT_LT:
		op = ">="
		notice = "less than"
	case vd.VALIDATOR_INT_LTE:
		op = ">"
		notice = "less than or equal to"
	}
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	b := `	if !(this.%v %v %d) {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must be %v '%v'. "))
	}
`
	bf := new(bytes.Buffer)
	fmt.Fprintln(bf, fmt.Sprintf(b, fieldName, op, vi, fieldName, fieldName, notice, vi))
	return bf.Bytes(), nil
}

// 正则
func (g *generator) regexValid(structName, fieldName, key, reg string) ([]byte, error) {
	_, err := regexp.Compile(reg)
	if err != nil {
		return nil, err
	}
	bf := new(bytes.Buffer)
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"

	regKey := fmt.Sprintf("_%v_%v", structName, fieldName)
	g.topVar[regKey] = fmt.Sprintf(`regexp.MustCompile("%v")`, reg)
	b := `	if %v.MatchString(this.%v){
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute %v must pass regular validation:'%v'. "))
	}

`
	fmt.Fprintln(bf, fmt.Sprintf(b, regKey, fieldName, fieldName, fieldName, reg))
	return bf.Bytes(), nil
}

// 解析tag，每个规则使用';'分隔，验证key和value使用':'分隔
func (g *generator) parseTag(tag string) ([]validator, error) {
	arr := strings.Split(tag, ";")

	var res []validator = make([]validator, len(arr))
	for i, v := range arr {
		kv := strings.Split(v, ":")
		if len(kv) == 0 {
			return nil, fmt.Errorf("验证规则key和 value数量错误")
		}
		if len(kv) == 1 {
			res[i] = validator{key: kv[0]}
		} else {
			res[i] = validator{key: kv[0], value: kv[1]}
		}

	}
	return res, nil
}
