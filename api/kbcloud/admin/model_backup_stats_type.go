// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BackupStatsType Totalsize and number of backups for the backup type
type BackupStatsType struct {
	// backup type
	Type *string `json:"type,omitempty"`
	// totalsize of backups for each engine, A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	Size *string `json:"size,omitempty"`
	// The number of backups for each engine
	Num *int64 `json:"num,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStatsType instantiates a new BackupStatsType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStatsType() *BackupStatsType {
	this := BackupStatsType{}
	return &this
}

// NewBackupStatsTypeWithDefaults instantiates a new BackupStatsType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsTypeWithDefaults() *BackupStatsType {
	this := BackupStatsType{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BackupStatsType) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BackupStatsType) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BackupStatsType) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BackupStatsType) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BackupStatsType) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *BackupStatsType) SetSize(v string) {
	o.Size = &v
}

// GetNum returns the Num field value if set, zero value otherwise.
func (o *BackupStatsType) GetNum() int64 {
	if o == nil || o.Num == nil {
		var ret int64
		return ret
	}
	return *o.Num
}

// GetNumOk returns a tuple with the Num field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetNumOk() (*int64, bool) {
	if o == nil || o.Num == nil {
		return nil, false
	}
	return o.Num, true
}

// HasNum returns a boolean if a field has been set.
func (o *BackupStatsType) HasNum() bool {
	return o != nil && o.Num != nil
}

// SetNum gets a reference to the given int64 and assigns it to the Num field.
func (o *BackupStatsType) SetNum(v int64) {
	o.Num = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
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
func (o *BackupStatsType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *string `json:"type,omitempty"`
		Size *string `json:"size,omitempty"`
		Num  *int64  `json:"num,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "size", "num"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Size = all.Size
	o.Num = all.Num

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
