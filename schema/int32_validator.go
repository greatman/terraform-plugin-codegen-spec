// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int32Validators type defines Int32Validator types
type Int32Validators []Int32Validator

// CustomValidators returns CustomValidator for each Float64Validator.
func (v Int32Validators) CustomValidators() CustomValidators {
	var customValidators CustomValidators

	for _, validator := range v {
		customValidator := validator.Custom

		if customValidator == nil {
			continue
		}

		customValidators = append(customValidators, customValidator)
	}

	return customValidators
}

// Equal returns true if the given Int32Validators is the same
// length, and each of the Int32Validator entries is equal.
func (v Int32Validators) Equal(other Int32Validators) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	validators := v.CustomValidators()

	otherValidators := other.CustomValidators()

	if len(validators) != len(otherValidators) {
		return false
	}

	validators.Sort()

	otherValidators.Sort()

	for k, validator := range validators {
		if !validator.Equal(otherValidators[k]) {
			return false
		}
	}

	return true
}

// Int32Validator type defines type and function that provides validation
// functionality.
type Int32Validator struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Int32Validator equal.
func (v Int32Validator) Equal(other Int32Validator) bool {
	return v.Custom.Equal(other.Custom)
}
