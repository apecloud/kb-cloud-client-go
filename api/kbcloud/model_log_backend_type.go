// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LogBackendType Log query backend configured for this environment.
type LogBackendType string

// List of LogBackendType.
const (
	LogBackendTypeLoki         LogBackendType = "loki"
	LogBackendTypeVictoriaLogs LogBackendType = "victoria-logs"
)

var allowedLogBackendTypeEnumValues = []LogBackendType{
	LogBackendTypeLoki,
	LogBackendTypeVictoriaLogs,
}

// GetAllowedValues returns the list of possible values.
func (v *LogBackendType) GetAllowedValues() []LogBackendType {
	return allowedLogBackendTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogBackendType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogBackendType(value)
	return nil
}

// NewLogBackendTypeFromValue returns a pointer to a valid LogBackendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogBackendTypeFromValue(v string) (*LogBackendType, error) {
	ev := LogBackendType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogBackendType: valid values are %v", v, allowedLogBackendTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogBackendType) IsValid() bool {
	for _, existing := range allowedLogBackendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogBackendType value.
func (v LogBackendType) Ptr() *LogBackendType {
	return &v
}
