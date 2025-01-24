// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupMethodOption struct {
	Name string `json:"name"`
	// The compatible full backup method for incremental backup method
	CompatibleMethod *string               `json:"compatibleMethod,omitempty"`
	Description      *LocalizedDescription `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupMethodOption instantiates a new BackupMethodOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupMethodOption(name string) *BackupMethodOption {
	this := BackupMethodOption{}
	this.Name = name
	return &this
}

// NewBackupMethodOptionWithDefaults instantiates a new BackupMethodOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupMethodOptionWithDefaults() *BackupMethodOption {
	this := BackupMethodOption{}
	return &this
}

// GetName returns the Name field value.
func (o *BackupMethodOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupMethodOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *BackupMethodOption) SetName(v string) {
	o.Name = v
}

// GetCompatibleMethod returns the CompatibleMethod field value if set, zero value otherwise.
func (o *BackupMethodOption) GetCompatibleMethod() string {
	if o == nil || o.CompatibleMethod == nil {
		var ret string
		return ret
	}
	return *o.CompatibleMethod
}

// GetCompatibleMethodOk returns a tuple with the CompatibleMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupMethodOption) GetCompatibleMethodOk() (*string, bool) {
	if o == nil || o.CompatibleMethod == nil {
		return nil, false
	}
	return o.CompatibleMethod, true
}

// HasCompatibleMethod returns a boolean if a field has been set.
func (o *BackupMethodOption) HasCompatibleMethod() bool {
	return o != nil && o.CompatibleMethod != nil
}

// SetCompatibleMethod gets a reference to the given string and assigns it to the CompatibleMethod field.
func (o *BackupMethodOption) SetCompatibleMethod(v string) {
	o.CompatibleMethod = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BackupMethodOption) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupMethodOption) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BackupMethodOption) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *BackupMethodOption) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupMethodOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.CompatibleMethod != nil {
		toSerialize["compatibleMethod"] = o.CompatibleMethod
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupMethodOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string               `json:"name"`
		CompatibleMethod *string               `json:"compatibleMethod,omitempty"`
		Description      *LocalizedDescription `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "compatibleMethod", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.CompatibleMethod = all.CompatibleMethod
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
