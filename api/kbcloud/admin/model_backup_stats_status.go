// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

// BackupStatsStatus Number of backups for the status
type BackupStatsStatus struct {
	// Backup status
	Status *string `json:"status,omitempty"`
	// Number of backups for each status
	Num *int64 `json:"num,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStatsStatus instantiates a new BackupStatsStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStatsStatus() *BackupStatsStatus {
	this := BackupStatsStatus{}
	return &this
}

// NewBackupStatsStatusWithDefaults instantiates a new BackupStatsStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsStatusWithDefaults() *BackupStatsStatus {
	this := BackupStatsStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupStatsStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupStatsStatus) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupStatsStatus) SetStatus(v string) {
	o.Status = &v
}

// GetNum returns the Num field value if set, zero value otherwise.
func (o *BackupStatsStatus) GetNum() int64 {
	if o == nil || o.Num == nil {
		var ret int64
		return ret
	}
	return *o.Num
}

// GetNumOk returns a tuple with the Num field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsStatus) GetNumOk() (*int64, bool) {
	if o == nil || o.Num == nil {
		return nil, false
	}
	return o.Num, true
}

// HasNum returns a boolean if a field has been set.
func (o *BackupStatsStatus) HasNum() bool {
	return o != nil && o.Num != nil
}

// SetNum gets a reference to the given int64 and assigns it to the Num field.
func (o *BackupStatsStatus) SetNum(v int64) {
	o.Num = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Num != nil {
		toSerialize["num"] = o.Num
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupStatsStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *string `json:"status,omitempty"`
		Num    *int64  `json:"num,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "num"})
	} else {
		return err
	}
	o.Status = all.Status
	o.Num = all.Num

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
