// Package sub_pkg
// @Author: yinwei
// @File: sub_pkg.model
// @Version: 1.0.0
// @Date: 2024/12/3 23:01

package sub_pkg

type Model1 struct {
	A string `json:"a" easy_valid:"required"`
	B int    `json:"b" easy_c_valid:"A=str_eq=success|int_lt:12"`
}
