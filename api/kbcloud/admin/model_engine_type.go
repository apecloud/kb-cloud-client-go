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
	EngineTypeRdbms        EngineType = "RDBMS"
	EngineTypeSearchEngine EngineType = "search-engine"
	EngineTypeKeyValue     EngineType = "key-value"
	EngineTypeTimeSeries   EngineType = "time-series"
	EngineTypeStreaming    EngineType = "streaming"
	EngineTypeLlm          EngineType = "LLM"
	EngineTypeVector       EngineType = "vector"
	EngineTypeDocument     EngineType = "document"
	EngineTypeGraph        EngineType = "graph"
	EngineTypeOther        EngineType = "other"
)

var allowedEngineTypeEnumValues = []EngineType{
	EngineTypeRdbms,
	EngineTypeSearchEngine,
	EngineTypeKeyValue,
	EngineTypeTimeSeries,
	EngineTypeStreaming,
	EngineTypeLlm,
	EngineTypeVector,
	EngineTypeDocument,
	EngineTypeGraph,
	EngineTypeOther,
}

// GetAllowedValues returns the list of possible values.
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
