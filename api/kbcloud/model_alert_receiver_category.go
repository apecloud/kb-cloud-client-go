// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION AlertReceiverCategory
type AlertReceiverCategory string

// List of AlertReceiverCategory.
const (
	ALERTRECEIVERCATEGORY_FEISHU         AlertReceiverCategory = "feishu"
	ALERTRECEIVERCATEGORY_WECHAT         AlertReceiverCategory = "wechat"
	ALERTRECEIVERCATEGORY_DINGTALK       AlertReceiverCategory = "dingtalk"
	ALERTRECEIVERCATEGORY_RECEIVER_GROUP AlertReceiverCategory = "receiver-group"
	ALERTRECEIVERCATEGORY_WEBHOOK        AlertReceiverCategory = "webhook"
)

var allowedAlertReceiverCategoryEnumValues = []AlertReceiverCategory{
	ALERTRECEIVERCATEGORY_FEISHU,
	ALERTRECEIVERCATEGORY_WECHAT,
	ALERTRECEIVERCATEGORY_DINGTALK,
	ALERTRECEIVERCATEGORY_RECEIVER_GROUP,
	ALERTRECEIVERCATEGORY_WEBHOOK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AlertReceiverCategory) GetAllowedValues() []AlertReceiverCategory {
	return allowedAlertReceiverCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertReceiverCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertReceiverCategory(value)
	return nil
}

// NewAlertReceiverCategoryFromValue returns a pointer to a valid AlertReceiverCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertReceiverCategoryFromValue(v string) (*AlertReceiverCategory, error) {
	ev := AlertReceiverCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertReceiverCategory: valid values are %v", v, allowedAlertReceiverCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertReceiverCategory) IsValid() bool {
	for _, existing := range allowedAlertReceiverCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertReceiverCategory value.
func (v AlertReceiverCategory) Ptr() *AlertReceiverCategory {
	return &v
}
