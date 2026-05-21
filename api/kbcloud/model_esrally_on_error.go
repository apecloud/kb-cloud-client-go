// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EsrallyOnError Rally request error handling behavior
type EsrallyOnError string

// List of EsrallyOnError.
const (
	EsrallyOnErrorAbort    EsrallyOnError = "abort"
	EsrallyOnErrorContinue EsrallyOnError = "continue"
)

var allowedEsrallyOnErrorEnumValues = []EsrallyOnError{
	EsrallyOnErrorAbort,
	EsrallyOnErrorContinue,
}

// GetAllowedValues returns the list of possible values.
func (v *EsrallyOnError) GetAllowedValues() []EsrallyOnError {
	return allowedEsrallyOnErrorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EsrallyOnError) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EsrallyOnError(value)
	return nil
}

// NewEsrallyOnErrorFromValue returns a pointer to a valid EsrallyOnError
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEsrallyOnErrorFromValue(v string) (*EsrallyOnError, error) {
	ev := EsrallyOnError(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EsrallyOnError: valid values are %v", v, allowedEsrallyOnErrorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EsrallyOnError) IsValid() bool {
	for _, existing := range allowedEsrallyOnErrorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EsrallyOnError value.
func (v EsrallyOnError) Ptr() *EsrallyOnError {
	return &v
}
