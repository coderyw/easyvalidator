package model

type TestStruct struct {
	A   int     `json:"a" easy_valid:"required:true"`
	B   *string `json:"b" easy_valid:"required;msg_exists:true"`
	C   float64 `json:"c" easy_valid:"float_lt:12"`
	Ceq float32 `json:"ceq" easy_valid:"float_eq:22.2" `
	E   uint8   `json:"e" easy_valid:"int_gte:33"`
	Eeq int32   `json:"eeq" easy_valid:"int_eq:33"`
	F   fs      `json:"f"`
	G   *Fs     `json:"g"`
	HH  []Fs    `json:"hh" easy_valid:"len_lt:1"`
	HHS []*Fs   `json:"hhs" easy_valid:"len_eq:2"`
	Id  string  `json:"id" easy_valid:"uuid_ver:1;uuid_ver"`

	Rex string `json:"rex" easy_valid:"regex:abc"`

	CrossVad string `json:"cross_vad" easy_c_valid:"A=int_lt:12;int_gt:1|C=float_lt:0.3|CrossVad=required"`

	CrossString   string `json:"cross_string" easy_c_valid:"Id=regex:failed|required"`
	CrossIntGt    string `json:"cross_int_gt" easy_c_valid:"A=int_gt:1|CrossIntGt=required"`
	CrossIntgte   uint   `json:"cross_intgte" easy_c_valid:"A=int_gte:1|int_gt:2"`
	CrossIntLt    string `json:"cross_int_lt" easy_c_valid:"A=int_lt:1|CrossIntLt=required"`
	CrossIntLte   string `json:"cross_int_lte" easy_c_valid:"A=int_lte:1|CrossIntLte=required"`
	CrossIntEq    string `json:"cross_int_eq" easy_c_valid:"A=int_lte:1|CrossIntEq=required"`
	CrossRequire  string `json:"cross_require" easy_c_valid:"A=required|CrossRequire=required"`
	CrossFloatGt  string `json:"cross_float_gt" easy_c_valid:"C=float_gt:12.1|CrossFloatGt=required"`
	CrossFloatLt  string `json:"cross_float_lt" easy_c_valid:"C=float_lt:12.1|CrossFloatLt=required"`
	CrossFloatGte string `json:"cross_float_gte" easy_c_valid:"C=float_gte:12.1|CrossFloatGte=required"`
	CrossFloatLte string `json:"cross_float_lte" easy_c_valid:"C=float_lte:12.1|CrossFloatLte=required"`
	CrossFloatEq  string `json:"cross_float_eq" easy_c_valid:"C=float_eq:12.1|CrossFloatEq=required"`
	CrossLenGt    string `json:"cross_len_gt" easy_c_valid:"Id=len_gt:1|CrossLenGt=required"`
	CrossLenLt    string `json:"cross_len_lt" easy_c_valid:"Id=len_lt:1|CrossLenLt=required"`
	CrossLenEq    string `json:"cross_len_eq" easy_c_valid:"Id=len_eq:1|CrossLenEq=required"`
	CrossUuidVer  string `json:"cross_uuid_ver" easy_c_valid:"Id=uuid_ver|CrossUuidVer=required"`
	CrossStrEq    string `json:"cross_str_eq" easy_c_valid:"Id=str_eq:dwe|CrossStrEq=required"`
}

type fs struct {
	Bs int `json:"bs"`
}
type Fs struct {
	DDD string `json:"ddd"`
}
