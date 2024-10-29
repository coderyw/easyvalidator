package model

import (
	fmt "fmt"
	validator_helper "github.com/coderyw/easyvalidator/helper"
	"regexp"
)

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

	if !(this.E < 33) {
		return validator_helper.FieldError("E", fmt.Errorf("The value of attribute 'E' must be greater than or equal to '33'. "))
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

	return nil
}

var _regex_uuidver1 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")
var _regex_uuidver0 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")
var _TestStruct_Rex = regexp.MustCompile("abc")
