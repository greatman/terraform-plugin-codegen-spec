// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int32PlanModifiers type defines Int32PlanModifier types
type Int32PlanModifiers []Int32PlanModifier

// CustomPlanModifiers returns CustomPlanModifier for each Int32PlanModifier.
func (v Int32PlanModifiers) CustomPlanModifiers() CustomPlanModifiers {
	var customPlanModifiers CustomPlanModifiers

	for _, planModifier := range v {
		customPlanModifier := planModifier.Custom

		if customPlanModifier == nil {
			continue
		}

		customPlanModifiers = append(customPlanModifiers, customPlanModifier)
	}

	return customPlanModifiers
}

// Equal returns true if the given Int32PlanModifiers is the same
// length, and each of the Int32PlanModifier entries is equal.
func (v Int32PlanModifiers) Equal(other Int32PlanModifiers) bool {
	if v == nil && other == nil {
		return true
	}

	if v == nil || other == nil {
		return false
	}

	if len(v) != len(other) {
		return false
	}

	planModifiers := v.CustomPlanModifiers()

	otherPlanModifiers := other.CustomPlanModifiers()

	if len(planModifiers) != len(otherPlanModifiers) {
		return false
	}

	planModifiers.Sort()

	otherPlanModifiers.Sort()

	for k, planModifier := range planModifiers {
		if !planModifier.Equal(otherPlanModifiers[k]) {
			return false
		}
	}

	return true
}

// Int32PlanModifier type defines type and function that provides plan modification
// functionality.
type Int32PlanModifier struct {
	Custom *CustomPlanModifier `json:"custom,omitempty"`
}

// Equal returns true if the fields of the given Int32PlanModifier are equal.
func (v Int32PlanModifier) Equal(other Int32PlanModifier) bool {
	return v.Custom.Equal(other.Custom)
}
