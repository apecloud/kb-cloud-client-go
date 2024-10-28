// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NetworkChaosDirection Specifies the direction of network chaos effects. 'to' affects outgoing traffic, 'from' affects incoming traffic, and 'both' affects traffic in both directions.
type NetworkChaosDirection string

// List of NetworkChaosDirection.
const (
	NETWORKCHAOSDIRECTION_TO   NetworkChaosDirection = "to"
	NETWORKCHAOSDIRECTION_FROM NetworkChaosDirection = "from"
	NETWORKCHAOSDIRECTION_BOTH NetworkChaosDirection = "both"
)

var allowedNetworkChaosDirectionEnumValues = []NetworkChaosDirection{
	NETWORKCHAOSDIRECTION_TO,
	NETWORKCHAOSDIRECTION_FROM,
	NETWORKCHAOSDIRECTION_BOTH,
}

// GetAllowedValues returns the list of possible values.
func (v *NetworkChaosDirection) GetAllowedValues() []NetworkChaosDirection {
	return allowedNetworkChaosDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkChaosDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkChaosDirection(value)
	return nil
}

// NewNetworkChaosDirectionFromValue returns a pointer to a valid NetworkChaosDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkChaosDirectionFromValue(v string) (*NetworkChaosDirection, error) {
	ev := NetworkChaosDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkChaosDirection: valid values are %v", v, allowedNetworkChaosDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkChaosDirection) IsValid() bool {
	for _, existing := range allowedNetworkChaosDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkChaosDirection value.
func (v NetworkChaosDirection) Ptr() *NetworkChaosDirection {
	return &v
}
