// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOptionLicenseFile struct {
	// Target license file name written into the engine license Secret.
	Name string `json:"name"`
	// Stable upload field key used by clients when building a multi-file license bundle.
	Field       string                `json:"field"`
	Title       *LocalizedDescription `json:"title,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEngineOptionLicenseFile instantiates a new EngineOptionLicenseFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionLicenseFile(name string, field string) *EngineOptionLicenseFile {
	this := EngineOptionLicenseFile{}
	this.Name = name
	this.Field = field
	return &this
}

// NewEngineOptionLicenseFileWithDefaults instantiates a new EngineOptionLicenseFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionLicenseFileWithDefaults() *EngineOptionLicenseFile {
	this := EngineOptionLicenseFile{}
	return &this
}

// GetName returns the Name field value.
func (o *EngineOptionLicenseFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineOptionLicenseFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineOptionLicenseFile) SetName(v string) {
	o.Name = v
}

// GetField returns the Field field value.
func (o *EngineOptionLicenseFile) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *EngineOptionLicenseFile) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *EngineOptionLicenseFile) SetField(v string) {
	o.Field = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EngineOptionLicenseFile) GetTitle() LocalizedDescription {
	if o == nil || o.Title == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicenseFile) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EngineOptionLicenseFile) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given LocalizedDescription and assigns it to the Title field.
func (o *EngineOptionLicenseFile) SetTitle(v LocalizedDescription) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineOptionLicenseFile) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicenseFile) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineOptionLicenseFile) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *EngineOptionLicenseFile) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionLicenseFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["field"] = o.Field
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionLicenseFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string               `json:"name"`
		Field       *string               `json:"field"`
		Title       *LocalizedDescription `json:"title,omitempty"`
		Description *LocalizedDescription `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Field = *all.Field
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
