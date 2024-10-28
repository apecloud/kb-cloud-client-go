// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PodChaosAction specify the action to be performed
type PodChaosAction string

// List of PodChaosAction.
const (
	PODCHAOSACTION_POD_FAILURE PodChaosAction = "pod-failure"
	PODCHAOSACTION_POD_KILL    PodChaosAction = "pod-kill"
)

var allowedPodChaosActionEnumValues = []PodChaosAction{
	PODCHAOSACTION_POD_FAILURE,
	PODCHAOSACTION_POD_KILL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PodChaosAction) GetAllowedValues() []PodChaosAction {
	return allowedPodChaosActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PodChaosAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PodChaosAction(value)
	return nil
}

// NewPodChaosActionFromValue returns a pointer to a valid PodChaosAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPodChaosActionFromValue(v string) (*PodChaosAction, error) {
	ev := PodChaosAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PodChaosAction: valid values are %v", v, allowedPodChaosActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PodChaosAction) IsValid() bool {
	for _, existing := range allowedPodChaosActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PodChaosAction value.
func (v PodChaosAction) Ptr() *PodChaosAction {
	return &v
}
