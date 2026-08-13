// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseDiagnosticSourceStatus string

// List of KingbaseDiagnosticSourceStatus.
const (
	KingbaseDiagnosticSourceStatusSuccess KingbaseDiagnosticSourceStatus = "success"
	KingbaseDiagnosticSourceStatusPartial KingbaseDiagnosticSourceStatus = "partial"
	KingbaseDiagnosticSourceStatusFailed  KingbaseDiagnosticSourceStatus = "failed"
	KingbaseDiagnosticSourceStatusSkipped KingbaseDiagnosticSourceStatus = "skipped"
	KingbaseDiagnosticSourceStatusNoData  KingbaseDiagnosticSourceStatus = "noData"
)

var allowedKingbaseDiagnosticSourceStatusEnumValues = []KingbaseDiagnosticSourceStatus{
	KingbaseDiagnosticSourceStatusSuccess,
	KingbaseDiagnosticSourceStatusPartial,
	KingbaseDiagnosticSourceStatusFailed,
	KingbaseDiagnosticSourceStatusSkipped,
	KingbaseDiagnosticSourceStatusNoData,
}

// GetAllowedValues returns the list of possible values.
func (v *KingbaseDiagnosticSourceStatus) GetAllowedValues() []KingbaseDiagnosticSourceStatus {
	return allowedKingbaseDiagnosticSourceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *KingbaseDiagnosticSourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = KingbaseDiagnosticSourceStatus(value)
	return nil
}

// NewKingbaseDiagnosticSourceStatusFromValue returns a pointer to a valid KingbaseDiagnosticSourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewKingbaseDiagnosticSourceStatusFromValue(v string) (*KingbaseDiagnosticSourceStatus, error) {
	ev := KingbaseDiagnosticSourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for KingbaseDiagnosticSourceStatus: valid values are %v", v, allowedKingbaseDiagnosticSourceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v KingbaseDiagnosticSourceStatus) IsValid() bool {
	for _, existing := range allowedKingbaseDiagnosticSourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KingbaseDiagnosticSourceStatus value.
func (v KingbaseDiagnosticSourceStatus) Ptr() *KingbaseDiagnosticSourceStatus {
	return &v
}
