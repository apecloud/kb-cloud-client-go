// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterConnectionPoolMode Connection reuse mode. The MVP supports session pooling only.
type ClusterConnectionPoolMode string

// List of ClusterConnectionPoolMode.
const (
	ClusterConnectionPoolModeSession ClusterConnectionPoolMode = "session"
)

var allowedClusterConnectionPoolModeEnumValues = []ClusterConnectionPoolMode{
	ClusterConnectionPoolModeSession,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterConnectionPoolMode) GetAllowedValues() []ClusterConnectionPoolMode {
	return allowedClusterConnectionPoolModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterConnectionPoolMode) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterConnectionPoolMode(value)
	return nil
}

// NewClusterConnectionPoolModeFromValue returns a pointer to a valid ClusterConnectionPoolMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterConnectionPoolModeFromValue(v string) (*ClusterConnectionPoolMode, error) {
	ev := ClusterConnectionPoolMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterConnectionPoolMode: valid values are %v", v, allowedClusterConnectionPoolModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterConnectionPoolMode) IsValid() bool {
	for _, existing := range allowedClusterConnectionPoolModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterConnectionPoolMode value.
func (v ClusterConnectionPoolMode) Ptr() *ClusterConnectionPoolMode {
	return &v
}
