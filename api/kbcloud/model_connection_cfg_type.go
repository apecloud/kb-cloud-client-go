// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ConnectionCfgType string

// List of ConnectionCfgType.
const (
	ConnectionCfgTypeDatabase ConnectionCfgType = "database"
)

var allowedConnectionCfgTypeEnumValues = []ConnectionCfgType{
	ConnectionCfgTypeDatabase,
}

// GetAllowedValues returns the list of possible values.
func (v *ConnectionCfgType) GetAllowedValues() []ConnectionCfgType {
	return allowedConnectionCfgTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConnectionCfgType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConnectionCfgType(value)
	return nil
}

// NewConnectionCfgTypeFromValue returns a pointer to a valid ConnectionCfgType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConnectionCfgTypeFromValue(v string) (*ConnectionCfgType, error) {
	ev := ConnectionCfgType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConnectionCfgType: valid values are %v", v, allowedConnectionCfgTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConnectionCfgType) IsValid() bool {
	for _, existing := range allowedConnectionCfgTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionCfgType value.
func (v ConnectionCfgType) Ptr() *ConnectionCfgType {
	return &v
}
