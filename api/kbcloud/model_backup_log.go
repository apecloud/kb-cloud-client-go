// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

// BackupLog backup workload logs
type BackupLog struct {
	// items is the list of backupLogByPod objects
	Items []BackupLogByPod `json:"items,omitempty"`
	// backup id
	BackupId *string `json:"backupId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupLog instantiates a new BackupLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupLog() *BackupLog {
	this := BackupLog{}
	return &this
}

// NewBackupLogWithDefaults instantiates a new BackupLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupLogWithDefaults() *BackupLog {
	this := BackupLog{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BackupLog) GetItems() []BackupLogByPod {
	if o == nil || o.Items == nil {
		var ret []BackupLogByPod
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupLog) GetItemsOk() (*[]BackupLogByPod, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BackupLog) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []BackupLogByPod and assigns it to the Items field.
func (o *BackupLog) SetItems(v []BackupLogByPod) {
	o.Items = v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *BackupLog) GetBackupId() string {
	if o == nil || o.BackupId == nil {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupLog) GetBackupIdOk() (*string, bool) {
	if o == nil || o.BackupId == nil {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *BackupLog) HasBackupId() bool {
	return o != nil && o.BackupId != nil
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *BackupLog) SetBackupId(v string) {
	o.BackupId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.BackupId != nil {
		toSerialize["backupId"] = o.BackupId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items    []BackupLogByPod `json:"items,omitempty"`
		BackupId *string          `json:"backupId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "backupId"})
	} else {
		return err
	}
	o.Items = all.Items
	o.BackupId = all.BackupId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
