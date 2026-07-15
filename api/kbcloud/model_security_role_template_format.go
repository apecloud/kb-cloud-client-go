// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SecurityRoleTemplateFormat string

// List of SecurityRoleTemplateFormat.
const (
	SecurityRoleTemplateFormatString SecurityRoleTemplateFormat = "string"
	SecurityRoleTemplateFormatJson   SecurityRoleTemplateFormat = "json"
)

var allowedSecurityRoleTemplateFormatEnumValues = []SecurityRoleTemplateFormat{
	SecurityRoleTemplateFormatString,
	SecurityRoleTemplateFormatJson,
}

// GetAllowedValues returns the list of possible values.
func (v *SecurityRoleTemplateFormat) GetAllowedValues() []SecurityRoleTemplateFormat {
	return allowedSecurityRoleTemplateFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityRoleTemplateFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityRoleTemplateFormat(value)
	return nil
}

// NewSecurityRoleTemplateFormatFromValue returns a pointer to a valid SecurityRoleTemplateFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityRoleTemplateFormatFromValue(v string) (*SecurityRoleTemplateFormat, error) {
	ev := SecurityRoleTemplateFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityRoleTemplateFormat: valid values are %v", v, allowedSecurityRoleTemplateFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityRoleTemplateFormat) IsValid() bool {
	for _, existing := range allowedSecurityRoleTemplateFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityRoleTemplateFormat value.
func (v SecurityRoleTemplateFormat) Ptr() *SecurityRoleTemplateFormat {
	return &v
}
