// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// BackupStatsEngine Totalsize and number of backups for the engine 
type BackupStatsEngine struct {
	// Engine name
	EngineName *string `json:"engineName,omitempty"`
	// totalsize of backups for each engine, A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	BackupSize *string `json:"backupSize,omitempty"`
	// The number of backups for each engine
	BackupNum *int64 `json:"backupNum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
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
// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *BackupStatsEngine) SetEngineName(v string) {
	o.EngineName = &v
}


// GetBackupSize returns the BackupSize field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetBackupSize() string {
	if o == nil || o.BackupSize == nil {
		var ret string
		return ret
	}
	return *o.BackupSize
}

// GetBackupSizeOk returns a tuple with the BackupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetBackupSizeOk() (*string, bool) {
	if o == nil || o.BackupSize == nil {
		return nil, false
	}
	return o.BackupSize, true
}

// HasBackupSize returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasBackupSize() bool {
	return o != nil && o.BackupSize != nil
}

// SetBackupSize gets a reference to the given string and assigns it to the BackupSize field.
func (o *BackupStatsEngine) SetBackupSize(v string) {
	o.BackupSize = &v
}


// GetBackupNum returns the BackupNum field value if set, zero value otherwise.
func (o *BackupStatsEngine) GetBackupNum() int64 {
	if o == nil || o.BackupNum == nil {
		var ret int64
		return ret
	}
	return *o.BackupNum
}

// GetBackupNumOk returns a tuple with the BackupNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsEngine) GetBackupNumOk() (*int64, bool) {
	if o == nil || o.BackupNum == nil {
		return nil, false
	}
	return o.BackupNum, true
}

// HasBackupNum returns a boolean if a field has been set.
func (o *BackupStatsEngine) HasBackupNum() bool {
	return o != nil && o.BackupNum != nil
}

// SetBackupNum gets a reference to the given int64 and assigns it to the BackupNum field.
func (o *BackupStatsEngine) SetBackupNum(v int64) {
	o.BackupNum = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsEngine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
	}
	if o.BackupSize != nil {
		toSerialize["backupSize"] = o.BackupSize
	}
	if o.BackupNum != nil {
		toSerialize["backupNum"] = o.BackupNum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupStatsEngine) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string `json:"engineName,omitempty"`
		BackupSize *string `json:"backupSize,omitempty"`
		BackupNum *int64 `json:"backupNum,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "engineName", "backupSize", "backupNum",  })
	} else {
		return err
	}
	o.EngineName = all.EngineName
	o.BackupSize = all.BackupSize
	o.BackupNum = all.BackupNum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
