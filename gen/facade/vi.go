// Package facade
// @Author: yinwei
// @File: vi
// @Version: 1.0.0
// @Date: 2024/3/13 16:31

package facade

type EasyMapInter interface {
	UnMarshalMapInterface(m map[string]interface{}) error
}

type EasyMapString interface {
	UnMarshalMap(m map[string]string) error
}

type EasyMap interface {
	EasyMapString
	EasyMapInter
}
