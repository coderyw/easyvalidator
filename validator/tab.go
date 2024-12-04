// Package main
// @Author: yinwei
// @File: tab
// @Version: 1.0.0
// @Date: 2024/10/28 16:10

package validator

const TAG_VALIDATOR = "easy_valid"

type UsagesValidator struct {
	Key string
	Zh  string
	En  string
}

var ValidatorUsages = []UsagesValidator{
	{
		Key: "regex",
		Zh:  "正则表达式",
		En:  "Uses a Golang RE2-syntax regex to match the field contents",
	},
	{
		Key: "int_gt",
		Zh:  "属性值大于这个值",
		En:  "Field value of integer strictly greater than this value.",
	},
	{
		Key: "int_gte",
		Zh:  "属性值大于等于这个值",
		En:  "Field value of integer greater than or equal to this value.",
	}, {
		Key: "int_lt",
		Zh:  "属性值小于这个值",
		En:  "Field value of integer strictly smaller than this value.",
	}, {
		Key: "int_eq",
		Zh:  "属性值等于这个值",
		En:  "Field value of integer strictly equal to this value.",
	},
	{
		Key: "msg_exists",
		Zh:  "用于指针类型，值必须存在",
		En:  "Used for nested ptr types, requires that the message type exists.",
	}, {
		Key: "required",
		Zh:  "必填参数，参数非零值",
		En:  "The parameter is not a zero value in Golang language.",
	}, {
		Key: "float_gt",
		Zh:  "浮点数属性值大于这个值",
		En:  "Floating point attribute value greater than this value.",
	}, {
		Key: "float_lt",
		Zh:  "浮点数属性值小于这个值",
		En:  "Field value of double strictly smaller than this value.",
	}, {
		Key: "float_gte",
		Zh:  "浮点数属性值大于等于这个值",
		En:  "Floating-point value compared to which the field content should be greater or equal.",
	}, {
		Key: "float_lte",
		Zh:  "浮点数属性值小于等于这个值",
		En:  "Floating-point value compared to which the field content should be smaller or equal.",
	}, {
		Key: "float_eq",
		Zh:  "浮点数属性值等于这个值",
		En:  "Floating-point value compared to which the field content should be equal.",
	}, {
		Key: "len_gt",
		Zh:  "属性值的长度大于这个值，支持字符串和数组切片",
		En:  "Field value of length greater than this value. Support String and slice",
	}, {
		Key: "len_lt",
		Zh:  "属性值的长度小于这个值，支持字符串和数组切片",
		En:  "Field value of length less than this value. Support String and slice",
	}, {
		Key: "len_eq",
		Zh:  "属性值的长度等于这个值，支持字符串和数组切片",
		En:  "Field value of length equal tothis value. Support String and slice",
	}, {
		Key: "uuid_ver",
		Zh:  "验证字符串是否为UUID格式，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。",
		En:  "Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.",
	}, {
		Key: "str_eq",
		Zh:  "属性值等于这个值",
		En:  "Field value of string strictly equal to this value.",
	},
}

var ValidatorCrossUsages = []UsagesValidator{
	{
		Key: "regex",
		Zh:  "对应属性满足正则表达式时进行本属性校验",
		En:  "Uses a Golang RE2-syntax regex to match the field contents",
	},
	{
		Key: "int_gt",
		Zh:  "对应属性值大于这个值时进行本属性校验",
		En:  "Field value of integer strictly greater than this value.",
	},
	{
		Key: "int_gte",
		Zh:  "对应属性值大于等于这个值时进行本属性校验",
		En:  "Field value of integer greater than or equal to this value.",
	}, {
		Key: "int_lt",
		Zh:  "对应属性值小于这个值时进行本属性校验",
		En:  "Field value of integer strictly smaller than this value.",
	}, {
		Key: "int_eq",
		Zh:  "对应属性值等于这个值时进行本属性校验",
		En:  "Field value of integer strictly equal to this value.",
	},
	{
		Key: "required",
		Zh:  "对应属性存在时进行本属性校验",
		En:  "The parameter is not a zero value in Golang language.",
	}, {
		Key: "float_gt",
		Zh:  "对应浮点数属性值大于这个值时进行本属性校验",
		En:  "Floating point attribute value greater than this value.",
	}, {
		Key: "float_lt",
		Zh:  "对应浮点数属性值小于这个值时进行本属性校验",
		En:  "Field value of double strictly smaller than this value.",
	}, {
		Key: "float_gte",
		Zh:  "对应浮点数属性值大于等于这个值时进行本属性校验",
		En:  "Floating-point value compared to which the field content should be greater or equal.",
	}, {
		Key: "float_lte",
		Zh:  "对应浮点数属性值小于等于这个值时进行本属性校验",
		En:  "Floating-point value compared to which the field content should be smaller or equal.",
	}, {
		Key: "float_eq",
		Zh:  "对应浮点数属性值等于这个值时进行本属性校验",
		En:  "Floating-point value compared to which the field content should be equal.",
	}, {
		Key: "len_gt",
		Zh:  "对应属性值的长度大于这个值(支持字符串和数组切片)时进行本属性校验",
		En:  "Field value of length greater than this value. Support String and slice",
	}, {
		Key: "len_lt",
		Zh:  "对应属性值的长度小于这个值时进行本属性校验,支持字符串和数组切片",
		En:  "Field value of length less than this value. Support String and slice",
	}, {
		Key: "len_eq",
		Zh:  "对应属性值的长度等于这个值时进行本属性校验，支持字符串和数组切片",
		En:  "Field value of length equal tothis value. Support String and slice",
	}, {
		Key: "uuid_ver",
		Zh:  "对应属性字符串为UUID格式时进行本属性校验，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。",
		En:  "Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.",
	}, {
		Key: "str_eq",
		Zh:  "对应属性值等于这个值时进行本属性校验",
		En:  "Field value of string strictly equal to this value.",
	},
}
