// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OracleTablespaceExpandAction string

// List of OracleTablespaceExpandAction.
const (
	OracleTablespaceExpandActionResize  OracleTablespaceExpandAction = "resize"
	OracleTablespaceExpandActionAddFile OracleTablespaceExpandAction = "addFile"
)

var allowedOracleTablespaceExpandActionEnumValues = []OracleTablespaceExpandAction{
	OracleTablespaceExpandActionResize,
	OracleTablespaceExpandActionAddFile,
}

// GetAllowedValues returns the list of possible values.
func (v *OracleTablespaceExpandAction) GetAllowedValues() []OracleTablespaceExpandAction {
	return allowedOracleTablespaceExpandActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OracleTablespaceExpandAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OracleTablespaceExpandAction(value)
	return nil
}

// NewOracleTablespaceExpandActionFromValue returns a pointer to a valid OracleTablespaceExpandAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOracleTablespaceExpandActionFromValue(v string) (*OracleTablespaceExpandAction, error) {
	ev := OracleTablespaceExpandAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OracleTablespaceExpandAction: valid values are %v", v, allowedOracleTablespaceExpandActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OracleTablespaceExpandAction) IsValid() bool {
	for _, existing := range allowedOracleTablespaceExpandActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OracleTablespaceExpandAction value.
func (v OracleTablespaceExpandAction) Ptr() *OracleTablespaceExpandAction {
	return &v
}
