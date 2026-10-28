// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlacementScope UI placement input scope for this mode. Omitted defaults to cluster. cluster displays AZ and NodeGroup selection for the cluster. component hides cluster placement and displays selection for each component, or for each instance template when the component declares templates. This does not restrict API placement overrides.
type PlacementScope string

// List of PlacementScope.
const (
	PlacementScopeCluster   PlacementScope = "cluster"
	PlacementScopeComponent PlacementScope = "component"
)

var allowedPlacementScopeEnumValues = []PlacementScope{
	PlacementScopeCluster,
	PlacementScopeComponent,
}

// GetAllowedValues returns the list of possible values.
func (v *PlacementScope) GetAllowedValues() []PlacementScope {
	return allowedPlacementScopeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PlacementScope) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PlacementScope(value)
	return nil
}

// NewPlacementScopeFromValue returns a pointer to a valid PlacementScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlacementScopeFromValue(v string) (*PlacementScope, error) {
	ev := PlacementScope(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PlacementScope: valid values are %v", v, allowedPlacementScopeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlacementScope) IsValid() bool {
	for _, existing := range allowedPlacementScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlacementScope value.
func (v PlacementScope) Ptr() *PlacementScope {
	return &v
}
