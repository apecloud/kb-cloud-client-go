// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceTemplateAttribute Heterogeneous attribute the frontend may render on an instance template at create.
type InstanceTemplateAttribute string

// List of InstanceTemplateAttribute.
const (
	InstanceTemplateAttributeStorageClassName InstanceTemplateAttribute = "storageClassName"
	InstanceTemplateAttributeAvailabilityZone InstanceTemplateAttribute = "availabilityZone"
	InstanceTemplateAttributeEnv              InstanceTemplateAttribute = "env"
	InstanceTemplateAttributeAnnotations      InstanceTemplateAttribute = "annotations"
	InstanceTemplateAttributeLabels           InstanceTemplateAttribute = "labels"
	InstanceTemplateAttributeClassCode        InstanceTemplateAttribute = "classCode"
)

var allowedInstanceTemplateAttributeEnumValues = []InstanceTemplateAttribute{
	InstanceTemplateAttributeStorageClassName,
	InstanceTemplateAttributeAvailabilityZone,
	InstanceTemplateAttributeEnv,
	InstanceTemplateAttributeAnnotations,
	InstanceTemplateAttributeLabels,
	InstanceTemplateAttributeClassCode,
}

// GetAllowedValues returns the list of possible values.
func (v *InstanceTemplateAttribute) GetAllowedValues() []InstanceTemplateAttribute {
	return allowedInstanceTemplateAttributeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InstanceTemplateAttribute) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InstanceTemplateAttribute(value)
	return nil
}

// NewInstanceTemplateAttributeFromValue returns a pointer to a valid InstanceTemplateAttribute
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInstanceTemplateAttributeFromValue(v string) (*InstanceTemplateAttribute, error) {
	ev := InstanceTemplateAttribute(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InstanceTemplateAttribute: valid values are %v", v, allowedInstanceTemplateAttributeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InstanceTemplateAttribute) IsValid() bool {
	for _, existing := range allowedInstanceTemplateAttributeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceTemplateAttribute value.
func (v InstanceTemplateAttribute) Ptr() *InstanceTemplateAttribute {
	return &v
}
