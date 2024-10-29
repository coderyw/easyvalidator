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
		return validator_helper.FieldError("A", fmt.Errorf("value '%v' must required", this.A))
	}
	if nil == this.B {
		return validator_helper.FieldError("B", fmt.Errorf("value '%v' must required", this.B))
	}
	if nil == this.B {
		return validator_helper.FieldError("msg_exists", fmt.Errorf("value '%v' must exist", this.B))
	}
	if this.B != nil {
		if err := validator_helper.CallValidatorIfExists(this.B); err != nil {
			return validator_helper.FieldError("B", err)
		}
	}
	if !(this.C >= 12) {
		return validator_helper.FieldError("C", fmt.Errorf("value '%v' must be less than '12'", this.C))
	}

	if !(this.E < 33) {
		return validator_helper.FieldError("E", fmt.Errorf("value '%v' must be greater than or equal to '33'", this.E))
	}

	if len(this.HH) >= 1 {
		return validator_helper.FieldError("HH", fmt.Errorf("value '%v' must be less than '1'", this.HH))
	}
	if len(this.HHS) != 2 {
		return validator_helper.FieldError("HHS", fmt.Errorf("value '%v' must be equal to '2'", this.HHS))
	}
	if _regex_uuidver1.MatchString(this.Id) {
		return validator_helper.FieldError("Id", fmt.Errorf("value '%v' must be a correct uuid with version '1'", this.Id))
	}
	if _regex_uuidver0.MatchString(this.Id) {
		return validator_helper.FieldError("Id", fmt.Errorf("value '%v' must be a correct uuid with version '0'", this.Id))
	}

	return nil
}

func (this *Fs) Validate() error {

	return nil
}

var _regex_uuidver1 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")
var _regex_uuidver0 = regexp.MustCompile("^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$")