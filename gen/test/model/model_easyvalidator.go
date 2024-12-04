package model

import (
	fmt "fmt"
	validator_helper "github.com/coderyw/easyvalidator/helper"
	"regexp"
)

func (this *fs) Validate() error {

	return nil
}

func (this *TestStruct) Validate() error {
	if 0 == this.A {
		return validator_helper.FieldError("A", fmt.Errorf("The value of attribute 'A' must required. "))
	}
	if nil == this.B {
		return validator_helper.FieldError("B", fmt.Errorf("The value of attribute 'B' must required. "))
	}
	if nil == this.B {
		return validator_helper.FieldError("B", fmt.Errorf("The value of attribute 'B' must exist. "))
	}
	if this.B != nil {
		if err := validator_helper.CallValidatorIfExists(this.B); err != nil {
			return validator_helper.FieldError("B", err)
		}
	}
	if !(this.C >= 12) {
		return validator_helper.FieldError("C", fmt.Errorf("The value of attribute 'C' must be less than '12'. "))
	}

	if !(this.Ceq == 22.2) {
		return validator_helper.FieldError("Ceq", fmt.Errorf("The value of attribute 'Ceq' must be equal to '22.2'. "))
	}

	if !(this.E < 33) {
		return validator_helper.FieldError("E", fmt.Errorf("The value of attribute 'E' must be greater than or equal to '33'. "))
	}

	if !(this.Eeq == 33) {
		return validator_helper.FieldError("Eeq", fmt.Errorf("The value of attribute 'Eeq' must be equal to '33'. "))
	}

	if len(this.HH) >= 1 {
		return validator_helper.FieldError("HH", fmt.Errorf("The length of the value of attribute HH must be less than 1. "))
	}
	if len(this.HHS) != 2 {
		return validator_helper.FieldError("HHS", fmt.Errorf("The length of the value of attribute HHS must be equal to 2. "))
	}
	if _regex_uuidver1.MatchString(this.Id) {
		return validator_helper.FieldError("Id", fmt.Errorf("The value of attribute Id must conform to the format of UUID version 1. "))
	}
	if _regex_uuidver0.MatchString(this.Id) {
		return validator_helper.FieldError("Id", fmt.Errorf("The value of attribute Id must conform to the format of UUID version 0. "))
	}
	if _TestStruct_Rex.MatchString(this.Rex) {
		return validator_helper.FieldError("Rex", fmt.Errorf("The value of attribute Rex must pass regular validation:'abc'. "))
	}

	if this.A < 12 && this.A > 1 && this.C < 0.3 {
		if len(this.CrossVad) == 0 {
			return validator_helper.FieldError("CrossVad", fmt.Errorf("When 'A' less than '12' and 'A' greater than '1' and 'C' less than '0.3', the value of attribute 'CrossVad' must required. "))
		}
	}
	if _Id_Id.MatchString(this.Id) {
		if len(this.CrossString) == 0 {
			return validator_helper.FieldError("CrossString", fmt.Errorf("When The value of attribute 'Id' pass regular validation:'failed', the value of attribute 'CrossString' must required. "))
		}
	}
	if this.A > 1 {
		if len(this.CrossIntGt) == 0 {
			return validator_helper.FieldError("CrossIntGt", fmt.Errorf("When 'A' greater than '1', the value of attribute 'CrossIntGt' must required. "))
		}
	}
	if this.A >= 1 {
		if !(this.CrossIntgte <= 2) {
			return validator_helper.FieldError("CrossIntgte", fmt.Errorf("When 'A' greater than or equal to '1', the value of attribute 'CrossIntgte' must be greater than '2'. "))
		}

	}
	if this.A < 1 {
		if len(this.CrossIntLt) == 0 {
			return validator_helper.FieldError("CrossIntLt", fmt.Errorf("When 'A' less than '1', the value of attribute 'CrossIntLt' must required. "))
		}
	}
	if this.A <= 1 {
		if len(this.CrossIntLte) == 0 {
			return validator_helper.FieldError("CrossIntLte", fmt.Errorf("When 'A' less than or equal to '1', the value of attribute 'CrossIntLte' must required. "))
		}
	}
	if this.A <= 1 {
		if len(this.CrossIntEq) == 0 {
			return validator_helper.FieldError("CrossIntEq", fmt.Errorf("When 'A' less than or equal to '1', the value of attribute 'CrossIntEq' must required. "))
		}
	}
	if 0 != this.A {
		if len(this.CrossRequire) == 0 {
			return validator_helper.FieldError("CrossRequire", fmt.Errorf("When 'A' is required, the value of attribute 'CrossRequire' must required. "))
		}
	}
	if this.C > 12.1 {
		if len(this.CrossFloatGt) == 0 {
			return validator_helper.FieldError("CrossFloatGt", fmt.Errorf("When 'C' greater than '12.1', the value of attribute 'CrossFloatGt' must required. "))
		}
	}
	if this.C < 12.1 {
		if len(this.CrossFloatLt) == 0 {
			return validator_helper.FieldError("CrossFloatLt", fmt.Errorf("When 'C' less than '12.1', the value of attribute 'CrossFloatLt' must required. "))
		}
	}
	if this.C >= 12.1 {
		if len(this.CrossFloatGte) == 0 {
			return validator_helper.FieldError("CrossFloatGte", fmt.Errorf("When 'C' greater than or equal to '12.1', the value of attribute 'CrossFloatGte' must required. "))
		}
	}
	if this.C <= 12.1 {
		if len(this.CrossFloatLte) == 0 {
			return validator_helper.FieldError("CrossFloatLte", fmt.Errorf("When 'C' less than or equal to '12.1', the value of attribute 'CrossFloatLte' must required. "))
		}
	}
	if this.C == 12.1 {
		if len(this.CrossFloatEq) == 0 {
			return validator_helper.FieldError("CrossFloatEq", fmt.Errorf("When 'C' equal to '12.1', the value of attribute 'CrossFloatEq' must required. "))
		}
	}
	if len(this.Id) > 1 {
		if len(this.CrossLenGt) == 0 {
			return validator_helper.FieldError("CrossLenGt", fmt.Errorf("When Id's length greater than '1', the value of attribute 'CrossLenGt' must required. "))
		}
	}
	if len(this.Id) < 1 {
		if len(this.CrossLenLt) == 0 {
			return validator_helper.FieldError("CrossLenLt", fmt.Errorf("When Id's length less than '1', the value of attribute 'CrossLenLt' must required. "))
		}
	}
	if len(this.Id) == 1 {
		if len(this.CrossLenEq) == 0 {
			return validator_helper.FieldError("CrossLenEq", fmt.Errorf("When Id's length equal to '1', the value of attribute 'CrossLenEq' must required. "))
		}
	}
	if _regex_uuidver0.MatchString(this.Id) {
		if len(this.CrossUuidVer) == 0 {
			return validator_helper.FieldError("CrossUuidVer", fmt.Errorf("When 'Id' conforms to the uuid format, the value of attribute 'CrossUuidVer' must required. "))
		}
	}
	if this.Id == "dwe" {
		if len(this.CrossStrEq) == 0 {
			return validator_helper.FieldError("CrossStrEq", fmt.Errorf("When 'Id' equal to 'dwe', the value of attribute 'CrossStrEq' must required. "))
		}
	}

	return nil
}

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

func (this *Fs) Validate() error {

	return nil
}

var _TestStruct_Rex = regexp.MustCompile("abc")
var _Id_Id = regexp.MustCompile("failed")
var _regex_uuidver1 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")
var _regex_uuidver0 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")
