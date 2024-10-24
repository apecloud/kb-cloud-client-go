// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsExposeType Specifies the type of exposure for the KubeBlocks cluster.
type OpsExposeType string

// List of OpsExposeType.
const (
	OPSEXPOSETYPE_VPC      OpsExposeType = "vpc"
	OPSEXPOSETYPE_INTERNET OpsExposeType = "internet"
)

var allowedOpsExposeTypeEnumValues = []OpsExposeType{
	OPSEXPOSETYPE_VPC,
	OPSEXPOSETYPE_INTERNET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpsExposeType) GetAllowedValues() []OpsExposeType {
	return allowedOpsExposeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsExposeType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsExposeType(value)
	return nil
}

// NewOpsExposeTypeFromValue returns a pointer to a valid OpsExposeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsExposeTypeFromValue(v string) (*OpsExposeType, error) {
	ev := OpsExposeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsExposeType: valid values are %v", v, allowedOpsExposeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsExposeType) IsValid() bool {
	for _, existing := range allowedOpsExposeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsExposeType value.
func (v OpsExposeType) Ptr() *OpsExposeType {
	return &v
}
