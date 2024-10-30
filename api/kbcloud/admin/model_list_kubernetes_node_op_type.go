// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ListKubernetesNodeOpType operation for list nodes, either `in` or `notin`
type ListKubernetesNodeOpType string

// List of ListKubernetesNodeOpType.
const (
	ListKubernetesNodeOpTypeIn ListKubernetesNodeOpType = "in"
	ListKubernetesNodeOpTypeNotin ListKubernetesNodeOpType = "notin"
)

var allowedListKubernetesNodeOpTypeEnumValues = []ListKubernetesNodeOpType{
	ListKubernetesNodeOpTypeIn,
	ListKubernetesNodeOpTypeNotin,
}

// GetAllowedValues returns the list of possible values.
func (v *ListKubernetesNodeOpType) GetAllowedValues() []ListKubernetesNodeOpType {
	return allowedListKubernetesNodeOpTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListKubernetesNodeOpType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListKubernetesNodeOpType(value)
	return nil
}

// NewListKubernetesNodeOpTypeFromValue returns a pointer to a valid ListKubernetesNodeOpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListKubernetesNodeOpTypeFromValue(v string) (*ListKubernetesNodeOpType, error) {
	ev := ListKubernetesNodeOpType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListKubernetesNodeOpType: valid values are %v", v, allowedListKubernetesNodeOpTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListKubernetesNodeOpType) IsValid() bool {
	for _, existing := range allowedListKubernetesNodeOpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListKubernetesNodeOpType value.
func (v ListKubernetesNodeOpType) Ptr() *ListKubernetesNodeOpType {
	return &v
}
