// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsAuditRisk string

// List of DmsAuditRisk.
const (
	DmsAuditRiskUnknown  DmsAuditRisk = "unknown"
	DmsAuditRiskReadOnly DmsAuditRisk = "read_only"
	DmsAuditRiskDdl      DmsAuditRisk = "ddl"
	DmsAuditRiskDml      DmsAuditRisk = "dml"
	DmsAuditRiskHigh     DmsAuditRisk = "high"
)

var allowedDmsAuditRiskEnumValues = []DmsAuditRisk{
	DmsAuditRiskUnknown,
	DmsAuditRiskReadOnly,
	DmsAuditRiskDdl,
	DmsAuditRiskDml,
	DmsAuditRiskHigh,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsAuditRisk) GetAllowedValues() []DmsAuditRisk {
	return allowedDmsAuditRiskEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsAuditRisk) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsAuditRisk(value)
	return nil
}

// NewDmsAuditRiskFromValue returns a pointer to a valid DmsAuditRisk
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsAuditRiskFromValue(v string) (*DmsAuditRisk, error) {
	ev := DmsAuditRisk(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsAuditRisk: valid values are %v", v, allowedDmsAuditRiskEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsAuditRisk) IsValid() bool {
	for _, existing := range allowedDmsAuditRiskEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsAuditRisk value.
func (v DmsAuditRisk) Ptr() *DmsAuditRisk {
	return &v
}
