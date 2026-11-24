// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OracleTablespaceExpand One DDL operation per request. Use resize with fileId and sizeMB, or addFile with file. Shrinking and adding files to BIGFILE tablespaces are rejected.
type OracleTablespaceExpand struct {
	Action OracleTablespaceExpandAction `json:"action"`
	FileId *int64                       `json:"fileId,omitempty"`
	SizeMb *int64                       `json:"sizeMB,omitempty"`
	File   *OracleTablespaceFileSpec    `json:"file,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOracleTablespaceExpand instantiates a new OracleTablespaceExpand object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOracleTablespaceExpand(action OracleTablespaceExpandAction) *OracleTablespaceExpand {
	this := OracleTablespaceExpand{}
	this.Action = action
	return &this
}

// NewOracleTablespaceExpandWithDefaults instantiates a new OracleTablespaceExpand object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOracleTablespaceExpandWithDefaults() *OracleTablespaceExpand {
	this := OracleTablespaceExpand{}
	return &this
}

// GetAction returns the Action field value.
func (o *OracleTablespaceExpand) GetAction() OracleTablespaceExpandAction {
	if o == nil {
		var ret OracleTablespaceExpandAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OracleTablespaceExpand) GetActionOk() (*OracleTablespaceExpandAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *OracleTablespaceExpand) SetAction(v OracleTablespaceExpandAction) {
	o.Action = v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *OracleTablespaceExpand) GetFileId() int64 {
	if o == nil || o.FileId == nil {
		var ret int64
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceExpand) GetFileIdOk() (*int64, bool) {
	if o == nil || o.FileId == nil {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *OracleTablespaceExpand) HasFileId() bool {
	return o != nil && o.FileId != nil
}

// SetFileId gets a reference to the given int64 and assigns it to the FileId field.
func (o *OracleTablespaceExpand) SetFileId(v int64) {
	o.FileId = &v
}

// GetSizeMb returns the SizeMb field value if set, zero value otherwise.
func (o *OracleTablespaceExpand) GetSizeMb() int64 {
	if o == nil || o.SizeMb == nil {
		var ret int64
		return ret
	}
	return *o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceExpand) GetSizeMbOk() (*int64, bool) {
	if o == nil || o.SizeMb == nil {
		return nil, false
	}
	return o.SizeMb, true
}

// HasSizeMb returns a boolean if a field has been set.
func (o *OracleTablespaceExpand) HasSizeMb() bool {
	return o != nil && o.SizeMb != nil
}

// SetSizeMb gets a reference to the given int64 and assigns it to the SizeMb field.
func (o *OracleTablespaceExpand) SetSizeMb(v int64) {
	o.SizeMb = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *OracleTablespaceExpand) GetFile() OracleTablespaceFileSpec {
	if o == nil || o.File == nil {
		var ret OracleTablespaceFileSpec
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleTablespaceExpand) GetFileOk() (*OracleTablespaceFileSpec, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *OracleTablespaceExpand) HasFile() bool {
	return o != nil && o.File != nil
}

// SetFile gets a reference to the given OracleTablespaceFileSpec and assigns it to the File field.
func (o *OracleTablespaceExpand) SetFile(v OracleTablespaceFileSpec) {
	o.File = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OracleTablespaceExpand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.FileId != nil {
		toSerialize["fileId"] = o.FileId
	}
	if o.SizeMb != nil {
		toSerialize["sizeMB"] = o.SizeMb
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OracleTablespaceExpand) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action *OracleTablespaceExpandAction `json:"action"`
		FileId *int64                        `json:"fileId,omitempty"`
		SizeMb *int64                        `json:"sizeMB,omitempty"`
		File   *OracleTablespaceFileSpec     `json:"file,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"action", "fileId", "sizeMB", "file"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.FileId = all.FileId
	o.SizeMb = all.SizeMb
	if all.File != nil && all.File.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.File = all.File

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
