// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NetworkChaosAction Specifies the type of network chaos action to be performed. Options include introducing delay, packet loss, duplication, corruption, network partition, or bandwidth limitation.
type NetworkChaosAction string

// List of NetworkChaosAction.
const (
	NETWORKCHAOSACTION_DELAY     NetworkChaosAction = "delay"
	NETWORKCHAOSACTION_LOSS      NetworkChaosAction = "loss"
	NETWORKCHAOSACTION_DUPLICATE NetworkChaosAction = "duplicate"
	NETWORKCHAOSACTION_CORRUPT   NetworkChaosAction = "corrupt"
	NETWORKCHAOSACTION_PARTITION NetworkChaosAction = "partition"
	NETWORKCHAOSACTION_BANDWIDTH NetworkChaosAction = "bandwidth"
)

var allowedNetworkChaosActionEnumValues = []NetworkChaosAction{
	NETWORKCHAOSACTION_DELAY,
	NETWORKCHAOSACTION_LOSS,
	NETWORKCHAOSACTION_DUPLICATE,
	NETWORKCHAOSACTION_CORRUPT,
	NETWORKCHAOSACTION_PARTITION,
	NETWORKCHAOSACTION_BANDWIDTH,
}

// GetAllowedValues returns the list of possible values.
func (v *NetworkChaosAction) GetAllowedValues() []NetworkChaosAction {
	return allowedNetworkChaosActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NetworkChaosAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NetworkChaosAction(value)
	return nil
}

// NewNetworkChaosActionFromValue returns a pointer to a valid NetworkChaosAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkChaosActionFromValue(v string) (*NetworkChaosAction, error) {
	ev := NetworkChaosAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkChaosAction: valid values are %v", v, allowedNetworkChaosActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkChaosAction) IsValid() bool {
	for _, existing := range allowedNetworkChaosActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkChaosAction value.
func (v NetworkChaosAction) Ptr() *NetworkChaosAction {
	return &v
}
