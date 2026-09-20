// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateOverlayField Overlay field the frontend may render on an instance template.
type InstanceTemplateOverlayField string

// List of InstanceTemplateOverlayField.
const (
	InstanceTemplateOverlayFieldStorageClass     InstanceTemplateOverlayField = "storageClass"
	InstanceTemplateOverlayFieldAvailabilityZone InstanceTemplateOverlayField = "availabilityZone"
	InstanceTemplateOverlayFieldNodeGroup        InstanceTemplateOverlayField = "nodeGroup"
	InstanceTemplateOverlayFieldEnv              InstanceTemplateOverlayField = "env"
	InstanceTemplateOverlayFieldAnnotations      InstanceTemplateOverlayField = "annotations"
)

var allowedInstanceTemplateOverlayFieldEnumValues = []InstanceTemplateOverlayField{
	InstanceTemplateOverlayFieldStorageClass,
	InstanceTemplateOverlayFieldAvailabilityZone,
	InstanceTemplateOverlayFieldNodeGroup,
	InstanceTemplateOverlayFieldEnv,
	InstanceTemplateOverlayFieldAnnotations,
}

// GetAllowedValues returns the list of possible values.
func (v *InstanceTemplateOverlayField) GetAllowedValues() []InstanceTemplateOverlayField {
	return allowedInstanceTemplateOverlayFieldEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InstanceTemplateOverlayField) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InstanceTemplateOverlayField(value)
	return nil
}

// NewInstanceTemplateOverlayFieldFromValue returns a pointer to a valid InstanceTemplateOverlayField
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInstanceTemplateOverlayFieldFromValue(v string) (*InstanceTemplateOverlayField, error) {
	ev := InstanceTemplateOverlayField(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InstanceTemplateOverlayField: valid values are %v", v, allowedInstanceTemplateOverlayFieldEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InstanceTemplateOverlayField) IsValid() bool {
	for _, existing := range allowedInstanceTemplateOverlayFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceTemplateOverlayField value.
func (v InstanceTemplateOverlayField) Ptr() *InstanceTemplateOverlayField {
	return &v
}
