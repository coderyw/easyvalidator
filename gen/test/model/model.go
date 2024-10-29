package model

type TestStruct struct {
	A   int     `json:"a" easy_valid:"required:true"`
	B   *string `json:"b" easy_valid:"required;msg_exists:true"`
	C   float64 `json:"c" easy_valid:"float_lt:12"`
	E   uint8   `json:"e" easy_valid:"int_gte:33"`
	F   fs      `json:"f"`
	G   *Fs     `json:"g"`
	HH  []Fs    `json:"hh" easy_valid:"len_lt:1"`
	HHS []*Fs   `json:"hhs" easy_valid:"len_eq:2"`
	Id  string  `json:"id" easy_valid:"uuid_ver:1;uuid_ver"`
}

type fs struct {
	Bs int `json:"bs"`
}
type Fs struct {
	DDD string `json:"ddd"`
}
