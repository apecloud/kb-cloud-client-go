// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BackupStatsEngine Totalsize and number of backups for the engine
type BackupStatsEngine struct {
	// Engine name
	Engine *string `json:"engine,omitempty"`
	// totalsize of backups for each engine, A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	Size *string `json:"size,omitempty"`
	// The number of backups for each engine
	Num *int64 `json:"num,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStatsEngine instantiates a new BackupStatsEngine object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStatsEngine() *BackupStatsEngine {
	this := BackupStatsEngine{}
	return &this
}

// NewBackupStatsEngineWithDefaults instantiates a new BackupStatsEngine object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsEngineWithDefaults() *BackupStatsEngine {
	this := BackupStatsEngine{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *BackupStatsEngine) SetEngine(v string) {
	o.Engine = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *BackupStatsEngine) SetSize(v string) {
	o.Size = &v
}

// GetNum returns the Num field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetNum() int64 {
	if o == nil || o.Num == nil {
		var ret int64
		return ret
	}
	return *o.Num
}

// GetNumOk returns a tuple with the Num field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetNumOk() (*int64, bool) {
	if o == nil || o.Num == nil {
		return nil, false
	}
	return o.Num, true
}

// HasNum returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasNum() bool {
	return o != nil && o.Num != nil
}

// SetNum gets a reference to the given int64 and assigns it to the Num field.
func (o *BackupStatsEngine) SetNum(v int64) {
	o.Num = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
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
func (o *BackupStatsEngine) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine *string `json:"engine,omitempty"`
		Size   *string `json:"size,omitempty"`
		Num    *int64  `json:"num,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "size", "num"})
	} else {
		return err
	}
	o.Engine = all.Engine
	o.Size = all.Size
	o.Num = all.Num

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
