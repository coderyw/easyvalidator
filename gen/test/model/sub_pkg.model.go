// Package sub_pkg
// @Author: yinwei
// @File: sub_pkg.model
// @Version: 1.0.0
// @Date: 2024/12/3 23:01

package model

type Model1 struct {
	A string  `json:"a" easy_valid:"required"`
	C float32 `json:"c" easy_valid:"float_gt:12.1"`
	B int     `json:"b" easy_c_valid:"A=str_eq:success|int_lt:12"`
	D int     `json:"d" easy_c_valid:"A=str_eq:success|D=int_lt:12"`
}
