// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateOp Operation that can be declared on instanceTemplate.ops.
type InstanceTemplateOp string

// List of InstanceTemplateOp.
const (
	InstanceTemplateOpHscale InstanceTemplateOp = "hscale"
)

var allowedInstanceTemplateOpEnumValues = []InstanceTemplateOp{
	InstanceTemplateOpHscale,
}

// GetAllowedValues returns the list of possible values.
func (v *InstanceTemplateOp) GetAllowedValues() []InstanceTemplateOp {
	return allowedInstanceTemplateOpEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InstanceTemplateOp) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InstanceTemplateOp(value)
	return nil
}

// NewInstanceTemplateOpFromValue returns a pointer to a valid InstanceTemplateOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInstanceTemplateOpFromValue(v string) (*InstanceTemplateOp, error) {
	ev := InstanceTemplateOp(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InstanceTemplateOp: valid values are %v", v, allowedInstanceTemplateOpEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InstanceTemplateOp) IsValid() bool {
	for _, existing := range allowedInstanceTemplateOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceTemplateOp value.
func (v InstanceTemplateOp) Ptr() *InstanceTemplateOp {
	return &v
}
