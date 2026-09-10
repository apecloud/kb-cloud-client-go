// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchShardState string

// List of ElasticsearchShardState.
const (
	ElasticsearchShardStateStarted      ElasticsearchShardState = "STARTED"
	ElasticsearchShardStateInitializing ElasticsearchShardState = "INITIALIZING"
	ElasticsearchShardStateRelocating   ElasticsearchShardState = "RELOCATING"
	ElasticsearchShardStateUnassigned   ElasticsearchShardState = "UNASSIGNED"
)

var allowedElasticsearchShardStateEnumValues = []ElasticsearchShardState{
	ElasticsearchShardStateStarted,
	ElasticsearchShardStateInitializing,
	ElasticsearchShardStateRelocating,
	ElasticsearchShardStateUnassigned,
}

// GetAllowedValues returns the list of possible values.
func (v *ElasticsearchShardState) GetAllowedValues() []ElasticsearchShardState {
	return allowedElasticsearchShardStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ElasticsearchShardState) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ElasticsearchShardState(value)
	return nil
}

// NewElasticsearchShardStateFromValue returns a pointer to a valid ElasticsearchShardState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewElasticsearchShardStateFromValue(v string) (*ElasticsearchShardState, error) {
	ev := ElasticsearchShardState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ElasticsearchShardState: valid values are %v", v, allowedElasticsearchShardStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ElasticsearchShardState) IsValid() bool {
	for _, existing := range allowedElasticsearchShardStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ElasticsearchShardState value.
func (v ElasticsearchShardState) Ptr() *ElasticsearchShardState {
	return &v
}
