# EasyValidator

这是一个高效的参数校验工具。

在方便开发和执行效率方面，选择了效率。基于定义的tag，执行插件生成对应的校验方法。

## 安装方式

```go
go install github.com/coderyw/easyvalidator
```

## 使用方式

```shell
# 编译model下的所有go文件，然后所有同包内的go文件生成的validator全在一个文件内
easyvalidator -all ./model
```

## 参数说明

- -all
          generate validator for all structs in a file
- -del_before
          开始处理前删除以前存在的文件 (default true)
- -exclude_suff string
          排除文件后缀名
- -include_suff string
          只需要包含这个后缀的go文件
- -leave_temps
          do not delete temporary files
- -noformat
          do not run 'gofmt -w' on output file
- -output_filename string
          specify the filename of the output
- -pkg
          process the whole package instead of just the given file
- -usages
          print usage information

## 支持的校验规则

```shell
# 执行下面命令，输出所有支持的校验规则
easyvalidator -usages

# 输出结果
Usage of validator key:
tag name: easy_valid
example: 

type UserInfo struct { 
        Name string `json:"name" easy_valid:"required"` 
}

单个属性校验(Explain of single validator key):
        regex
                正则表达式
                Uses a Golang RE2-syntax regex to match the field contents
        int_gt
                属性值大于这个值
                Field value of integer strictly greater than this value.
        int_gte
                属性值大于等于这个值
                Field value of integer greater than or equal to this value.
        int_lt
                属性值小于这个值
                Field value of integer strictly smaller than this value.
        int_eq
                属性值等于这个值
                Field value of integer strictly equal to this value.
        msg_exists
                用于指针类型，值必须存在
                Used for nested ptr types, requires that the message type exists.
        required
                必填参数，参数非零值
                The parameter is not a zero value in Golang language.
        float_gt
                浮点数属性值大于这个值
                Floating point attribute value greater than this value.
        float_lt
                浮点数属性值小于这个值
                Field value of double strictly smaller than this value.
        float_gte
                浮点数属性值大于等于这个值
                Floating-point value compared to which the field content should be greater or equal.
        float_lte
                浮点数属性值小于等于这个值
                Floating-point value compared to which the field content should be smaller or equal.
        float_eq
                浮点数属性值等于这个值
                Floating-point value compared to which the field content should be equal.
        len_gt
                属性值的长度大于这个值，支持字符串和数组切片
                Field value of length greater than this value. Support String and slice
        len_lt
                属性值的长度小于这个值，支持字符串和数组切片
                Field value of length less than this value. Support String and slice
        len_eq
                属性值的长度等于这个值，支持字符串和数组切片
                Field value of length equal tothis value. Support String and slice
        uuid_ver
                验证字符串是否为UUID格式，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
                Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.
        str_eq
                属性值等于这个值
                Field value of string strictly equal to this value.

多个属性联合校验(Joint verification of multiple attributes):
使用方式：Age=int_gte:1,int_lt:0|str_eq:12
说明：要求属性0<Age<=1时，校验本属性字符串长度是否等于2
        regex
                对应属性满足正则表达式时进行本属性校验
                Uses a Golang RE2-syntax regex to match the field contents
        int_gt
                对应属性值大于这个值时进行本属性校验
                Field value of integer strictly greater than this value.
        int_gte
                对应属性值大于等于这个值时进行本属性校验
                Field value of integer greater than or equal to this value.
        int_lt
                对应属性值小于这个值时进行本属性校验
                Field value of integer strictly smaller than this value.
        int_eq
                对应属性值等于这个值时进行本属性校验
                Field value of integer strictly equal to this value.
        required
                对应属性存在时进行本属性校验
                The parameter is not a zero value in Golang language.
        float_gt
                对应浮点数属性值大于这个值时进行本属性校验
                Floating point attribute value greater than this value.
        float_lt
                对应浮点数属性值小于这个值时进行本属性校验
                Field value of double strictly smaller than this value.
        float_gte
                对应浮点数属性值大于等于这个值时进行本属性校验
                Floating-point value compared to which the field content should be greater or equal.
        float_lte
                对应浮点数属性值小于等于这个值时进行本属性校验
                Floating-point value compared to which the field content should be smaller or equal.
        float_eq
                对应浮点数属性值等于这个值时进行本属性校验
                Floating-point value compared to which the field content should be equal.
        len_gt
                对应属性值的长度大于这个值(支持字符串和数组切片)时进行本属性校验
                Field value of length greater than this value. Support String and slice
        len_lt
                对应属性值的长度小于这个值时进行本属性校验,支持字符串和数组切片
                Field value of length less than this value. Support String and slice
        len_eq
                对应属性值的长度等于这个值时进行本属性校验，支持字符串和数组切片
                Field value of length equal tothis value. Support String and slice
        uuid_ver
                对应属性字符串为UUID格式时进行本属性校验，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
                Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.
        str_eq
                对应属性值等于这个值时进行本属性校验
                Field value of string strictly equal to this value.


```

### 单属性校验

struct的tag关键字是：easy_valid

1. 字符串

   - **regex**
     校验这个属性是否符合正则表达式

     Uses a Golang RE2-syntax regex to match the field contents

   - **str_eq**
     属性值等于这个值
     Field value of string strictly equal to this value.

   -  **uuid_ver**
     验证字符串是否为UUID格式，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
     Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.

2. 整型

   - **int_gt**
     属性值大于这个值
     Field value of integer strictly greater than this value.

   - **int_gte**
     属性值大于等于这个值
     Field value of integer greater than or equal to this value.

   - **int_lt**
     属性值小于这个值
     Field value of integer strictly smaller than this value.

   - **int_eq**
     属性值等于这个值
     Field value of integer strictly equal to this value.

3. 指针

   - **msg_exists**
     用于指针类型，值必须存在
     Used for nested ptr types, requires that the message type exists.

4. 小数

   - **float_gt**
     浮点数属性值大于这个值
     Floating point attribute value greater than this value.

   - **float_lt**
     浮点数属性值小于这个值
     Field value of double strictly smaller than this value.        
   - **float_gte**
     浮点数属性值大于等于这个值
     Floating-point value compared to which the field content should be greater or equal.
   - **float_lte**
     浮点数属性值小于等于这个值
     Floating-point value compared to which the field content should be smaller or equal.
   - **float_eq**
     浮点数属性值等于这个值
     Floating-point value compared to which the field content should be equal.

5. 综合

   - **len_gt**
     属性值的长度大于这个值，支持字符串和数组切片
     Field value of length greater than this value. Support String and slice

   - **len_lt**
     属性值的长度小于这个值，支持字符串和数组切片
     Field value of length less than this value. Support String and slice
   - **len_eq**
     属性值的长度等于这个值，支持字符串和数组切片
     Field value of length equal tothis value. Support String and slice
   - **required**
     必填参数，参数非零值
     The parameter is not a zero value in Golang language.

### 跨属性校验(在某个属性满足条件时，本属性进行参数校验)

多个属性联合校验(Joint verification of multiple attributes):
使用方式：Age=int_gte:1,int_lt:0|str_eq:12
说明：要求属性0<Age<=1时，校验本属性字符串长度是否等于2

1. 字符串

    - **regex**

      对应属性满足正则表达式时进行本属性校验

      Uses a Golang RE2-syntax regex to match the field contents

    - **uuid_ver**
      对应属性字符串为UUID格式时进行本属性校验，uuid_ver指定有效的UUID版本。有效的版本为0-5。如果uuid_ver为0，则接受所有UUID版本。
      Ensures that a string value is in UUID format. uuid_ver specifies the valid UUID versions. Valid values are: 0-5. If uuid_ver is 0 all UUID versions are accepted.
    - **str_eq**
      对应属性值等于这个值时进行本属性校验
      Field value of string strictly equal to this value.

2. 整型

    - **int_gt**
      对应属性值大于这个值时进行本属性校验
      Field value of integer strictly greater than this value.
    - **int_gte**
      对应属性值大于等于这个值时进行本属性校验
      Field value of integer greater than or equal to this value.
    - **int_lt**
      对应属性值小于这个值时进行本属性校验
      Field value of integer strictly smaller than this value.
    - **int_eq**
      对应属性值等于这个值时进行本属性校验
      Field value of integer strictly equal to this value.

3. 浮点数

    - **float_gt**
      对应浮点数属性值大于这个值时进行本属性校验
      Floating point attribute value greater than this value.
    - **float_lt**
      对应浮点数属性值小于这个值时进行本属性校验
      Field value of double strictly smaller than this value.
    - **float_gte**
      对应浮点数属性值大于等于这个值时进行本属性校验
      Floating-point value compared to which the field content should be greater or equal.
    - **float_lte**
      对应浮点数属性值小于等于这个值时进行本属性校验
      Floating-point value compared to which the field content should be smaller or equal.
    - **float_eq**
      对应浮点数属性值等于这个值时进行本属性校验
      Floating-point value compared to which the field content should be equal.

4. 综合

    - **required**

      对应属性存在时进行本属性校验
      The parameter is not a zero value in Golang language.

    - **len_gt**
      对应属性值的长度大于这个值(支持字符串和数组切片)时进行本属性校验
      Field value of length greater than this value. Support String and slice
    - **len_lt**
      对应属性值的长度小于这个值时进行本属性校验,支持字符串和数组切片
      Field value of length less than this value. Support String and slice
    - **len_eq**
      对应属性值的长度等于这个值时进行本属性校验，支持字符串和数组切片
      Field value of length equal tothis value. Support String and slice

## 使用示例

1. 单属性校验：

   使用tag关键字: `easy_valid`,格式为: `easy_valid:"{validatorKey}:{value};{validatorKey1}:{value1}"`

2. 跨属性校验

​	使用tag关键字：`easy_c_valid`格式为:`easy_c_valid:"FieldName={validatorKey}:{value};{validatorKey1}:{value1}|FieldName1={validatorKey}:{value}|{validatorKey}:{value}"`

示例：

```go
type Model1 struct {
	// 要求A不能为空
	A string  `json:"a" easy_valid:"required"`
  // 要求浮点数C要大于12.1
	C float32 `json:"c" easy_valid:"float_gt:12.1"`
  // 要求整数B，在字符串A=‘success’时，B>12
  // 在使用时，其他属性列需要使用属性的名称,后面跟'=',然后就是其他属性列需要满足的条件。
  // 本属性可以设置属性名称也可以不设置属性名称
  B int     `json:"b" easy_c_valid:"A=str_eq:success|int_lt:12"`
  // 这种加上D=int_lt:12与上面的int_lt:12等价
  D int     `json:"d" easy_c_valid:"A=str_eq:success|D=int_lt:12"`
}
```

使用`easyvalidator -all ./model`进行编译：

```go
func (this *Model1) Validate() error {
	if len(this.A) == 0 {
		return validator_helper.FieldError("A", fmt.Errorf("The value of attribute 'A' must required. "))
	}
	if !(this.C <= 12.1) {
		return validator_helper.FieldError("C", fmt.Errorf("The value of attribute 'C' must be greater than '12.1'. "))
	}

	if this.A == "success" {
		if !(this.B >= 12) {
			return validator_helper.FieldError("B", fmt.Errorf("When 'A' equal to 'success', the value of attribute 'B' must be less than '12'. "))
		}

	}
	if this.A == "success" {
		if !(this.D >= 12) {
			return validator_helper.FieldError("D", fmt.Errorf("When 'A' equal to 'success', the value of attribute 'D' must be less than '12'. "))
		}

	}

	return nil
}

```

