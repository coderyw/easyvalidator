// Package main
// @Author: yinwei
// @File: validator
// @Version: 1.0.0
// @Date: 2024/10/28 15:21

package validator

const (
	// 正则验证  Uses a Golang RE2-syntax regex to match the field contents.
	VALIDATOR_REGEX = "regex"
	// 属性值大于这个值 Field value of integer strictly greater than this value.
	VALIDATOR_INT_GT = "int_gt"
	// 属性值大于等于这个值 Field value of integer greater than or equal to this value.
	VALIDATOR_INT_GTE = "int_gte"
	// 属性值小于这个值 Field value of integer strictly smaller than this value.
	VALIDATOR_INT_LT = "int_lt"
	// 属性值小于等于 Field value of integer smaller than or equal to this value.
	VALIDATOR_INT_LTE = "int_lte"
	// 属性值等于 Field value of integer equal to this value.
	VALIDATOR_INT_EQ = "int_eq"

	// 用于指针类型，值必须存在 Used for nested ptr types, requires that the message type exists.
	VALIDATOR_MSG_EXISTS = "msg_exists"
	// 必填参数，参数非零值
	VALIDATOR_REQUIRED = "required"
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	//
	// 浮点数属性值大于这个值
	VALIDATOR_FLOAT_GT = "float_gt"
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	//
	// 浮点数属性值小于这个值
	VALIDATOR_FLOAT_LT = "float_lt"

	// Floating-point value compared to which the field content should be greater or equal.
	//
	// 浮点数属性值大于等于这个值
	VALIDATOR_FLOAT_GTE = "float_gte"
	// Floating-point value compared to which the field content should be smaller or equal.
	//
	// 浮点数属性值小于等于这个值
	VALIDATOR_FLOAT_LTE = "float_lte"
	VALIDATOR_FLOAT_EQ  = "float_eq"

	// Used for string fields, requires the string to be not empty (i.e different from "").
	//
	// 字符串属性值非空
	//VALIDATOR_STRING_NOT_EMPTY = "string_not_empty"
	//// Array field with at least this number of elements.
	////
	//// 数组切片最小长度
	//VALIDATOR_ARRAY_MIN_LEN = "array_min_len"
	//// Array field with at most this number of elements.
	////
	//// 数组切片最大长度
	//VALIDATOR_ARRAY_MAX_LEN = "array_max_len"

	// Field value of length greater than this value.
	//
	// 属性值的长度大于这个值，支持字符串和数组切片
	VALIDATOR_LEN_GT = "len_gt"

	// Field value of length less than this value.
	//
	// 属性值的长度小于这个值，支持字符串和数组切片
	VALIDATOR_LEN_LT = "len_lt"

	// Field value of length equal tothis value.
	//
	// 属性值的长度等于这个值，支持字符串和数组切片
	VALIDATOR_LEN_EQ = "len_eq"
	// Ensures that a string value is in UUID format.
	// uuid_ver specifies the valid UUID versions. Valid values are: 0-5.
	// If uuid_ver is 0 all UUID versions are accepted.
	//
	// 验证字符串是否为UUID格式，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
	VALIDATOR_UUID_VER = "uuid_ver"

	VALIDATOR_STR_EQ = "str_eq"
)

const (
	// 正则验证  Uses a Golang RE2-syntax regex to match the field contents.
	CROSS_VALIDATOR_REGEX = "regex"
	// 属性值大于这个值 Field value of integer strictly greater than this value.
	CROSS_VALIDATOR_INT_GT = "int_gt"
	// 属性值大于等于这个值 Field value of integer greater than or equal to this value.
	CROSS_VALIDATOR_INT_GTE = "int_gte"
	// 属性值小于这个值 Field value of integer strictly smaller than this value.
	CROSS_VALIDATOR_INT_LT = "int_lt"
	// 属性值小于等于 Field value of integer smaller than or equal to this value.
	CROSS_VALIDATOR_INT_LTE = "int_lte"
	// 属性值等于 Field value of integer equal to this value.
	CROSS_VALIDATOR_INT_EQ = "int_eq"
	// 必填参数，参数非零值
	CROSS_VALIDATOR_REQUIRED = "required"
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	//
	// 浮点数属性值大于这个值
	CROSS_VALIDATOR_FLOAT_GT = "float_gt"
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	//
	// 浮点数属性值小于这个值
	CROSS_VALIDATOR_FLOAT_LT = "float_lt"
	// Floating-point value compared to which the field content should be greater or equal.
	//
	// 浮点数属性值大于等于这个值
	CROSS_VALIDATOR_FLOAT_GTE = "float_gte"
	// Floating-point value compared to which the field content should be smaller or equal.
	//
	// 浮点数属性值小于等于这个值
	CROSS_VALIDATOR_FLOAT_LTE = "float_lte"

	CROSS_VALIDATOR_FLOAT_EQ = "float_eq"

	// Field value of length greater than this value.
	//
	// 属性值的长度大于这个值，支持字符串和数组切片
	CROSS_VALIDATOR_LEN_GT = "len_gt"

	// Field value of length less than this value.
	//
	// 属性值的长度小于这个值，支持字符串和数组切片
	CROSS_VALIDATOR_LEN_LT = "len_lt"

	// Field value of length equal tothis value.
	//
	// 属性值的长度等于这个值，支持字符串和数组切片
	CROSS_VALIDATOR_LEN_EQ = "len_eq"
	// Ensures that a string value is in UUID format.
	// uuid_ver specifies the valid UUID versions. Valid values are: 0-5.
	// If uuid_ver is 0 all UUID versions are accepted.
	//
	// 验证字符串是否为UUID格式，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
	CROSS_VALIDATOR_UUID_VER = "uuid_ver"

	CROSS_VALIDATOR_STR_EQ = "str_eq"
)
