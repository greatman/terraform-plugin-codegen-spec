// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

// Int32Type is a representation of a 32-bit integer.
type Int32Type struct {
	// CustomType is a customization of the Int32Type.
	CustomType *CustomType `json:"custom_type,omitempty"`
}
