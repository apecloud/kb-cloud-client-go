// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcConfigResourceType string

// List of CdcConfigResourceType.
const (
	CdcConfigResourceTypeConfigMap CdcConfigResourceType = "ConfigMap"
	CdcConfigResourceTypeSecret    CdcConfigResourceType = "Secret"
)

var allowedCdcConfigResourceTypeEnumValues = []CdcConfigResourceType{
	CdcConfigResourceTypeConfigMap,
	CdcConfigResourceTypeSecret,
}

// GetAllowedValues returns the list of possible values.
func (v *CdcConfigResourceType) GetAllowedValues() []CdcConfigResourceType {
	return allowedCdcConfigResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CdcConfigResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CdcConfigResourceType(value)
	return nil
}

// NewCdcConfigResourceTypeFromValue returns a pointer to a valid CdcConfigResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCdcConfigResourceTypeFromValue(v string) (*CdcConfigResourceType, error) {
	ev := CdcConfigResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CdcConfigResourceType: valid values are %v", v, allowedCdcConfigResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CdcConfigResourceType) IsValid() bool {
	for _, existing := range allowedCdcConfigResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CdcConfigResourceType value.
func (v CdcConfigResourceType) Ptr() *CdcConfigResourceType {
	return &v
}
