// Package gen
// @Author: yinwei
// @File: cross_encode
// @Version: 1.0.0
// @Date: 2024/12/3 16:48

package gen

import (
	"bytes"
	"fmt"
	vd "github.com/coderyw/easyvalidator/validator"
	"reflect"
	"regexp"
	"strconv"
)

func (g *generator) uuidVerCross(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, string, error) {
	if fv.Kind() != reflect.String {
		return nil, "", fmt.Errorf("uuid var only support string！")
	}
	var (
		val int64
		err error
	)
	if value != "" {
		val, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, "", err
		}
	}

	switch val {
	case 0, 1, 2, 3, 4, 5:
	default:
		return nil, "", fmt.Errorf("uuid_ver’s version only support 0~5.")
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
	b := ` %v.MatchString(this.%v) `
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	fmt.Fprint(bf, fmt.Sprintf(b, valib, field.Name))
	return bf.Bytes(), fmt.Sprintf("'%v' conforms to the uuid format", field.Name), nil
}

func (g *generator) lenValidCross(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, string, error) {
	switch fv.Kind() {
	case reflect.Map, reflect.String, reflect.Array, reflect.Slice:

	default:
		return nil, "", fmt.Errorf("len validator only support map,string,array,slice")
	}
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, "", err
	}

	var (
		op, notice string
	)

	switch key {
	case vd.VALIDATOR_LEN_GT:
		op = ">"
		notice = "greater than"
	case vd.VALIDATOR_LEN_LT:
		op = "<"
		notice = "less than"
	case vd.VALIDATOR_LEN_EQ:
		op = "=="
		notice = "equal to"
	}

	bf := new(bytes.Buffer)
	b := ` len(this.%v) %v %v`
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	fmt.Fprint(bf, fmt.Sprintf(b, field.Name, op, val))
	return bf.Bytes(), fmt.Sprintf("%v's length %v '%v'", field.Name, notice, val), nil
}

func (g *generator) floatValidCross(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, string, error) {
	switch fv.Kind() {
	case reflect.Float64, reflect.Float32:
	default:
		return nil, "", fmt.Errorf("only support type float")
	}

	fieldName := field.Name

	vi, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, "", err
	}

	var (
		op     string
		notice string
	)
	switch key {
	case vd.VALIDATOR_FLOAT_GT:
		op = ">"
		notice = "greater than"
	case vd.VALIDATOR_FLOAT_GTE:
		op = ">="
		notice = "greater than or equal to"
	case vd.VALIDATOR_FLOAT_LT:
		op = "<"
		notice = "less than"
	case vd.VALIDATOR_FLOAT_LTE:
		op = "<="
		notice = "less than or equal to"
	case vd.CROSS_VALIDATOR_FLOAT_EQ:
		op = "=="
		notice = "equal to"
	default:
		return nil, "", fmt.Errorf("miss validator tag %v", key)
	}
	b := ` this.%v %v %v `
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	fmt.Fprint(bf, fmt.Sprintf(b, fieldName, op, vi))
	return bf.Bytes(), fmt.Sprintf("'%v' %v '%v'", field.Name, notice, vi), nil
}

func (g *generator) requiredValidCross(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, string, error) {
	if value == "false" || value == "0" {
		return nil, "", nil
	}
	var b string
	switch fv.Kind() {
	case reflect.Ptr:
		b = ` nil != this.%v `
		b = fmt.Sprintf(b, field.Name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64, reflect.Float32:
		b = ` 0 != this.%v `
		b = fmt.Sprintf(b, field.Name)
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		b = ` len(this.%v) !=0 `
		b = fmt.Sprintf(b, field.Name)
	default:
		return nil, "", fmt.Errorf("required only support int,int8,int16,int32,int64,uin,uin8,uin16,uin32,uin64,float32,float64,array,map,pointer,slice,string")
	}
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	bf := new(bytes.Buffer)
	fmt.Fprint(bf, b)
	return bf.Bytes(), fmt.Sprintf("'%v' is required", field.Name), nil
}

// int
func (g *generator) intValidCross(fv reflect.Type, field reflect.StructField, key, value string) ([]byte, string, error) {
	switch fv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	default:
		return nil, "", fmt.Errorf("only support type int")

	}
	fieldName := field.Name

	vi, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, "", err
	}

	var (
		op     string
		notice string
	)
	switch key {
	case vd.VALIDATOR_INT_GT:
		op = ">"
		notice = "greater than"
	case vd.VALIDATOR_INT_GTE:
		op = ">="
		notice = "greater than or equal to"
	case vd.VALIDATOR_INT_LT:
		op = "<"
		notice = "less than"
	case vd.VALIDATOR_INT_LTE:
		op = "<="
		notice = "less than or equal to"
	case vd.CROSS_VALIDATOR_INT_EQ:
		op = "=="
		notice = "equal to"
	default:
		return nil, "", fmt.Errorf("miss validator tag %v", key)
	}
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"
	b := ` this.%v %v %d `
	bf := new(bytes.Buffer)
	fmt.Fprint(bf, fmt.Sprintf(b, fieldName, op, vi))
	return bf.Bytes(), fmt.Sprintf("'%v' %v '%v'", field.Name, notice, vi), nil
}

// 正则
func (g *generator) regexValidCross(structName, fieldName, key, reg string) ([]byte, string, error) {
	_, err := regexp.Compile(reg)
	if err != nil {
		return nil, "", err
	}
	bf := new(bytes.Buffer)
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"

	regKey := fmt.Sprintf("_%v_%v", structName, fieldName)
	g.topVar[regKey] = fmt.Sprintf(`regexp.MustCompile("%v")`, reg)
	b := ` %v.MatchString(this.%v) `
	fmt.Fprint(bf, fmt.Sprintf(b, regKey, fieldName))
	return bf.Bytes(), fmt.Sprintf("The value of attribute '%v' pass regular validation:'%v'", fieldName, reg), nil
}

func (g *generator) strValidCross(fieldName, key, value string) ([]byte, string, error) {
	bf := new(bytes.Buffer)
	g.imports["validator_helper"] = "github.com/coderyw/easyvalidator/helper"
	g.imports["fmt"] = "fmt"

	b := fmt.Sprintf(` this.%v == "%v" `, fieldName, value)
	fmt.Fprint(bf, b)
	return bf.Bytes(), fmt.Sprintf("'%v' equal to '%v'", fieldName, value), nil
}
