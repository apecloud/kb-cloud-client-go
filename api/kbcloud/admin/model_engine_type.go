// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineType engine type
type EngineType string

// List of EngineType.
const (
	ENGINETYPE_RDBMS         EngineType = "RDBMS"
	ENGINETYPE_SEARCH_ENGINE EngineType = "search-engine"
	ENGINETYPE_KEY_VALUE     EngineType = "key-value"
	ENGINETYPE_TIME_SERIES   EngineType = "time-series"
	ENGINETYPE_STREAMING     EngineType = "streaming"
	ENGINETYPE_LLM           EngineType = "LLM"
	ENGINETYPE_VECTOR        EngineType = "vector"
	ENGINETYPE_DOCUMENT      EngineType = "document"
	ENGINETYPE_GRAPH         EngineType = "graph"
	ENGINETYPE_OTHER         EngineType = "other"
)

var allowedEngineTypeEnumValues = []EngineType{
	ENGINETYPE_RDBMS,
	ENGINETYPE_SEARCH_ENGINE,
	ENGINETYPE_KEY_VALUE,
	ENGINETYPE_TIME_SERIES,
	ENGINETYPE_STREAMING,
	ENGINETYPE_LLM,
	ENGINETYPE_VECTOR,
	ENGINETYPE_DOCUMENT,
	ENGINETYPE_GRAPH,
	ENGINETYPE_OTHER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EngineType) GetAllowedValues() []EngineType {
	return allowedEngineTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EngineType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EngineType(value)
	return nil
}

// NewEngineTypeFromValue returns a pointer to a valid EngineType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEngineTypeFromValue(v string) (*EngineType, error) {
	ev := EngineType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EngineType: valid values are %v", v, allowedEngineTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EngineType) IsValid() bool {
	for _, existing := range allowedEngineTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EngineType value.
func (v EngineType) Ptr() *EngineType {
	return &v
}
