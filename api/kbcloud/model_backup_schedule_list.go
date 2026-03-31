// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BackupScheduleList list of backup schedules
type BackupScheduleList struct {
	// items is the list of backup schedule objects
	Items []BackupSchedule `json:"items,omitempty"`
	// total count of backup schedules
	TotalCount *int64 `json:"totalCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupScheduleList instantiates a new BackupScheduleList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupScheduleList() *BackupScheduleList {
	this := BackupScheduleList{}
	return &this
}

// NewBackupScheduleListWithDefaults instantiates a new BackupScheduleList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupScheduleListWithDefaults() *BackupScheduleList {
	this := BackupScheduleList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BackupScheduleList) GetItems() []BackupSchedule {
	if o == nil || o.Items == nil {
		var ret []BackupSchedule
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleList) GetItemsOk() (*[]BackupSchedule, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BackupScheduleList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []BackupSchedule and assigns it to the Items field.
func (o *BackupScheduleList) SetItems(v []BackupSchedule) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BackupScheduleList) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleList) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BackupScheduleList) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *BackupScheduleList) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupScheduleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupScheduleList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      []BackupSchedule `json:"items,omitempty"`
		TotalCount *int64           `json:"totalCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "totalCount"})
	} else {
		return err
	}
	o.Items = all.Items
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
