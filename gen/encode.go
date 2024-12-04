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

// 单个交叉验证
type crossValidator struct {
	// 要交叉严重的field tag名称
	field string
	// 交叉验证的validator
	validator []validator
}

// 字段交叉验证
type crossValidatorGroup struct {
	// 字段的所有交叉验证
	crossValidators []crossValidator
	// 字段在瞒住交叉验证之后的结果
	validator []validator
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
		g.fieldMap[field.Name] = fieldreflect{
			field:        field,
			fieldReflect: field.Type,
		}
		fv = field.Type
		data, err := g.encodeField(t.Name(), fv, field, false)
		if err != nil {
			return err
		}
		if len(data) != 0 {
			buffer.Write(data)
		}

	}
	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		fv = field.Type
		data, err := g.encodeCrossField(t.Name(), fv, field, false)
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
		data, err := g.make(structName, fv, field, v, "")
		if err != nil {
			return nil, err
		}
		bf.Write(data)
	}
	return bf.Bytes(), nil
}

// 处理交叉验证
func (g *generator) encodeCrossField(structName string, fv reflect.Type, field reflect.StructField, isPtr bool) ([]byte, error) {
	vad := field.Tag.Get(TAG_CROSS_VALIDATOR)
	if len(vad) == 0 {
		return nil, nil
	}
	cv, err := g.parseCrossTag(field.Name, vad)
	if err != nil {
		return nil, err
	}
	if len(cv.crossValidators) == 0 || len(cv.validator) == 0 {
		return nil, nil
	}
	bf := new(bytes.Buffer)
	notify := strings.Builder{}
	notify.WriteString("When ")
	bf.WriteString("if ")
	// 遍历交叉的tag，生成对应的验证规则
	for i, v := range cv.crossValidators {
		for j, k := range v.validator {
			data, cmsg, err := g.makeCrossValidator(v.field, g.fieldMap[v.field].fieldReflect, g.fieldMap[v.field].field, k)
			if err != nil {
				return nil, err
			}

			if j != 0 || i != 0 {
				notify.WriteString(" and ")
				bf.WriteString("&&")
			}
			bf.Write(data)
			notify.WriteString(cmsg)
		}
	}

	bf.WriteString(" {\n")

	//遍历在交叉tag下自己的验证规则
	for _, v := range cv.validator {
		data, err := g.make(structName, fv, field, v, notify.String())
		if err != nil {
			return nil, err
		}
		bf.Write(data)
	}
	bf.WriteString("}\n")
	return bf.Bytes(), nil
}

func (g *generator) makeCrossValidator(structName string, fv reflect.Type, field reflect.StructField, valid validator) ([]byte, string, error) {
	name := field.Name
	switch valid.key {
	case vd.CROSS_VALIDATOR_REGEX:
		return g.regexValidCross(structName, name, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_INT_GT, vd.CROSS_VALIDATOR_INT_GTE, vd.CROSS_VALIDATOR_INT_LT, vd.CROSS_VALIDATOR_INT_LTE, vd.CROSS_VALIDATOR_INT_EQ:
		return g.intValidCross(fv, field, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_REQUIRED:
		return g.requiredValidCross(fv, field, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_FLOAT_GT, vd.CROSS_VALIDATOR_FLOAT_LT, vd.CROSS_VALIDATOR_FLOAT_GTE, vd.CROSS_VALIDATOR_FLOAT_LTE, vd.CROSS_VALIDATOR_FLOAT_EQ:
		return g.floatValidCross(fv, field, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_LEN_GT, vd.CROSS_VALIDATOR_LEN_LT, vd.CROSS_VALIDATOR_LEN_EQ:
		return g.lenValidCross(fv, field, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_UUID_VER:
		return g.uuidVerCross(fv, field, valid.key, valid.value)
	case vd.CROSS_VALIDATOR_STR_EQ:
		return g.strValidCross(field.Name, valid.key, valid.value)
	default:
		return nil, "", fmt.Errorf("错误的验证类型")
	}
}

func (g *generator) make(structName string, fv reflect.Type, field reflect.StructField, valid validator, crossMsg string) ([]byte, error) {
	name := field.Name
	switch valid.key {
	case vd.VALIDATOR_REGEX:
		return g.regexValid(structName, name, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_INT_GT, vd.VALIDATOR_INT_GTE, vd.VALIDATOR_INT_LT, vd.VALIDATOR_INT_LTE, vd.VALIDATOR_INT_EQ:
		return g.intValid(fv, field, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_MSG_EXISTS:
		return g.msgExistsValid(fv, field, valid.key, valid.value)
	case vd.VALIDATOR_REQUIRED:
		return g.requiredValid(fv, field, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_FLOAT_GT, vd.VALIDATOR_FLOAT_LT, vd.VALIDATOR_FLOAT_GTE, vd.VALIDATOR_FLOAT_LTE, vd.CROSS_VALIDATOR_FLOAT_EQ:
		return g.floatValid(fv, field, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_LEN_GT, vd.VALIDATOR_LEN_LT, vd.VALIDATOR_LEN_EQ:
		return g.lenValid(fv, field, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_UUID_VER:
		return g.uuidVer(fv, field, valid.key, valid.value, crossMsg)
	case vd.VALIDATOR_STR_EQ:
		return g.strValid(field.Name, valid.key, valid.value, crossMsg)
	default:
		return nil, fmt.Errorf("错误的验证类型")
	}
}

func (g *generator) uuidVer(fv reflect.Type, field reflect.StructField, key, value string, crossMsg string) ([]byte, error) {
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
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	b := ""
	if len(crossMsg) != 0 {
		b = `		if %v.MatchString(this.%v) {
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute %v must conform to the format of UUID version %v. "))
	}`
		b = fmt.Sprintf(b, valib, crossMsg, field.Name, field.Name, field.Name, val)
	} else {
		b = `	if %v.MatchString(this.%v) {
	return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute %v must conform to the format of UUID version %v. "))
}`
		b = fmt.Sprintf(b, valib, field.Name, field.Name, field.Name, val)
	}
	fmt.Fprintln(bf, b)
	return bf.Bytes(), nil
}

func (g *generator) lenValid(fv reflect.Type, field reflect.StructField, key, value string, crossMsg string) ([]byte, error) {
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

	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	b := ""
	if len(crossMsg) != 0 {
		b = `		if len(this.%v) %v %v {
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the length of the value of attribute %v must be %v %v. "))
	}`
		b = fmt.Sprintf(b, field.Name, op, val, crossMsg, field.Name, field.Name, notice, value)
	} else {
		b = `	if len(this.%v) %v %v{
	return validator_helper.FieldError("%v", fmt.Errorf("The length of the value of attribute %v must be %v %v. "))
}`
		b = fmt.Sprintf(b, field.Name, op, val, field.Name, field.Name, notice, value)
	}
	fmt.Fprintln(bf, b)

	return bf.Bytes(), nil
}

func (g *generator) floatValid(fv reflect.Type, field reflect.StructField, key, value, crossMsg string) ([]byte, error) {
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
	case vd.CROSS_VALIDATOR_FLOAT_EQ:
		op = "=="
		notice = "equal to"
	default:
		return nil, fmt.Errorf("miss validator tag %v", key)
	}

	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	b := ""
	if len(crossMsg) != 0 {
		b = `		if !(this.%v %v %v) {
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute '%v' must be %v '%v'. "))
	}
`
		b = fmt.Sprintf(b, fieldName, op, vi, fieldName, crossMsg, fieldName, notice, vi)
	} else {
		b = `	if !(this.%v %v %v) {
	return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must be %v '%v'. "))
}
`
		b = fmt.Sprintf(b, fieldName, op, vi, fieldName, fieldName, notice, vi)
	}
	fmt.Fprintln(bf, b)

	return bf.Bytes(), nil
}

func (g *generator) requiredValid(fv reflect.Type, field reflect.StructField, key, value, crossMsg string) ([]byte, error) {
	if value == "false" || value == "0" {
		return nil, nil
	}
	var b string

	var resp string
	if len(crossMsg) > 0 {
		resp = fmt.Sprintf(`	return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute '%v' must required. "))`, field.Name, crossMsg, field.Name)
	} else {
		resp = fmt.Sprintf(`return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must required. "))`, field.Name, field.Name)
	}

	switch fv.Kind() {
	case reflect.Ptr:
		b = `	if nil == this.%v{
	%v
	}`
		b = fmt.Sprintf(b, field.Name, resp)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64, reflect.Float32:
		b = `	if 0 == this.%v{
	%v
	}`
		b = fmt.Sprintf(b, field.Name, resp)
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		b = `	if len(this.%v) ==0{
	%v
	}`
		b = fmt.Sprintf(b, field.Name, resp)
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
func (g *generator) intValid(fv reflect.Type, field reflect.StructField, key, value, crossMsg string) ([]byte, error) {
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
	case vd.VALIDATOR_INT_EQ:
		op = "=="
		notice = "equal to"
	default:
		return nil, fmt.Errorf("miss validator tag %v", key)
	}
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	b := fmt.Sprintf(`	if !(this.%v %v %d) {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute '%v' must be %v '%v'. "))
	}
`, fieldName, op, vi, fieldName, fieldName, notice, vi)
	if len(crossMsg) > 0 {
		b = fmt.Sprintf(`	if !(this.%v %v %d) {
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute '%v' must be %v '%v'. "))
	}
`, fieldName, op, vi, fieldName, crossMsg, fieldName, notice, vi)
	}
	bf := new(bytes.Buffer)
	fmt.Fprintln(bf, b)
	return bf.Bytes(), nil
}

// 正则
func (g *generator) regexValid(structName, fieldName, key, reg, crossMsg string) ([]byte, error) {
	_, err := regexp.Compile(reg)
	if err != nil {
		return nil, err
	}
	bf := new(bytes.Buffer)
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"

	regKey := fmt.Sprintf("_%v_%v", structName, fieldName)
	g.topVar[regKey] = fmt.Sprintf(`regexp.MustCompile("%v")`, reg)
	b := fmt.Sprintf(`	if %v.MatchString(this.%v){
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute %v must pass regular validation:'%v'. "))
	}

`, regKey, fieldName, fieldName, fieldName, reg)
	if len(crossMsg) > 0 {
		b = fmt.Sprintf(`	if %v.MatchString(this.%v){
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute %v must pass regular validation:'%v'. "))
	}

`, regKey, fieldName, fieldName, fieldName, crossMsg, reg)
	}
	fmt.Fprintln(bf, b)
	return bf.Bytes(), nil
}
func (g *generator) strValid(fieldName, key, value, crossMsg string) ([]byte, error) {
	bf := new(bytes.Buffer)
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"

	b := fmt.Sprintf(`	if this.%v!= %v {
		return validator_helper.FieldError("%v", fmt.Errorf("The value of attribute %v must equal to:'%v'. "))
	}
`, fieldName, value, fieldName, fieldName, value)
	if len(crossMsg) > 0 {
		b = fmt.Sprintf(`	if this.%v!= %v{
		return validator_helper.FieldError("%v", fmt.Errorf("%v, the value of attribute %v must equal to:'%v'. "))
	}

`, fieldName, fieldName, fieldName, crossMsg, fieldName, value)
	}
	fmt.Fprintln(bf, b)
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

// 解析交叉验证
// CrossVad string `json:"cross_vad" easy_c_valid:"A=int_lt:12;int_gt:12|C=float_lt:0.3|CrossVad=required"`
func (g *generator) parseCrossTag(fieldName, tag string) (crossValidatorGroup, error) {
	arr := strings.Split(tag, "|")
	group := crossValidatorGroup{}

	// arr=[A=int_lt:12,int_gt:12 C=float_lt:0.3 CrossVad=required]
	for _, v := range arr {

		arr1 := strings.Split(v, "=")
		var wait string
		if len(arr1) == 1 {
			wait = arr1[0]
		} else {
			wait = arr1[1]
		}
		// arr1当前是 [A int_lt:12,int_gt:12]
		va, err := g.parseTag(wait)
		if err != nil {
			return group, err
		}
		if len(arr1) == 1 || arr1[0] == fieldName {
			group.validator = va
		} else {
			group.crossValidators = append(group.crossValidators, crossValidator{
				field:     arr1[0],
				validator: va,
			})
		}
	}
	return group, nil
}

func (g *generator) needReturn(needR []bool) bool {
	if len(needR) > 0 && needR[0] {
		return true
	}
	return false
}
