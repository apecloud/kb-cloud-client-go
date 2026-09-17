// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AiDataGatewayDataSourceStatus Defaults to enabled on creation; omitted on update preserves the current status.
type AiDataGatewayDataSourceStatus string

// List of AiDataGatewayDataSourceStatus.
const (
	AiDataGatewayDataSourceStatusEnabled  AiDataGatewayDataSourceStatus = "enabled"
	AiDataGatewayDataSourceStatusDisabled AiDataGatewayDataSourceStatus = "disabled"
)

var allowedAiDataGatewayDataSourceStatusEnumValues = []AiDataGatewayDataSourceStatus{
	AiDataGatewayDataSourceStatusEnabled,
	AiDataGatewayDataSourceStatusDisabled,
}

// GetAllowedValues returns the list of possible values.
func (v *AiDataGatewayDataSourceStatus) GetAllowedValues() []AiDataGatewayDataSourceStatus {
	return allowedAiDataGatewayDataSourceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiDataGatewayDataSourceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiDataGatewayDataSourceStatus(value)
	return nil
}

// NewAiDataGatewayDataSourceStatusFromValue returns a pointer to a valid AiDataGatewayDataSourceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiDataGatewayDataSourceStatusFromValue(v string) (*AiDataGatewayDataSourceStatus, error) {
	ev := AiDataGatewayDataSourceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiDataGatewayDataSourceStatus: valid values are %v", v, allowedAiDataGatewayDataSourceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiDataGatewayDataSourceStatus) IsValid() bool {
	for _, existing := range allowedAiDataGatewayDataSourceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiDataGatewayDataSourceStatus value.
func (v AiDataGatewayDataSourceStatus) Ptr() *AiDataGatewayDataSourceStatus {
	return &v
}
