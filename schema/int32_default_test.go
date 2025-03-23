// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/greatman/terraform-plugin-codegen-spec/schema"
)

func TestInt32Default_Equal(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		int32Default *schema.Int32Default
		other        *schema.Int32Default
		expected     bool
	}{
		"both_nil": {
			expected: true,
		},
		"int32_default_nil_other_not_nil": {
			other:    &schema.Int32Default{},
			expected: false,
		},
		"int32_default_static_nil_other_not_nil": {
			int32Default: &schema.Int32Default{},
			other: &schema.Int32Default{
				Static: pointer(int32(1234)),
			},
			expected: false,
		},
		"int32_default_static_not_nil_other_nil": {
			int32Default: &schema.Int32Default{
				Static: pointer(int32(1234)),
			},
			other:    &schema.Int32Default{},
			expected: false,
		},
		"match": {
			int32Default: &schema.Int32Default{
				Static: pointer(int32(1234)),
			},
			other: &schema.Int32Default{
				Static: pointer(int32(1234)),
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.int32Default.Equal(testCase.other)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
