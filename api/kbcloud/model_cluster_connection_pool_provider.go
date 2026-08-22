// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterConnectionPoolProvider Connection pool implementation managed by KBE.
type ClusterConnectionPoolProvider string

// List of ClusterConnectionPoolProvider.
const (
	ClusterConnectionPoolProviderPgbouncer ClusterConnectionPoolProvider = "pgbouncer"
)

var allowedClusterConnectionPoolProviderEnumValues = []ClusterConnectionPoolProvider{
	ClusterConnectionPoolProviderPgbouncer,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterConnectionPoolProvider) GetAllowedValues() []ClusterConnectionPoolProvider {
	return allowedClusterConnectionPoolProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterConnectionPoolProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterConnectionPoolProvider(value)
	return nil
}

// NewClusterConnectionPoolProviderFromValue returns a pointer to a valid ClusterConnectionPoolProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterConnectionPoolProviderFromValue(v string) (*ClusterConnectionPoolProvider, error) {
	ev := ClusterConnectionPoolProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterConnectionPoolProvider: valid values are %v", v, allowedClusterConnectionPoolProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterConnectionPoolProvider) IsValid() bool {
	for _, existing := range allowedClusterConnectionPoolProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterConnectionPoolProvider value.
func (v ClusterConnectionPoolProvider) Ptr() *ClusterConnectionPoolProvider {
	return &v
}
