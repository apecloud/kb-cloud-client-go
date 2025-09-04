// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CreateClassType struct {
	// Consists of lowercase English letters and numbers, 1-16 characters long. Used to add to the class type name.
	Name string `json:"name"`
	// The description of the class type. Optional.
	Description *string `json:"description,omitempty"`
	// The display name of the class type. Optional.
	DisplayName *string `json:"displayName,omitempty"`
	// The display name of the class type in English. Optional. If you switch the language to English, the display name will be displaynameEn.
	DisplayNameEn *string `json:"displayNameEn,omitempty"`
	// The description of the class type in English. Optional. If you switch the language to English, the description will be descriptionEn.
	DescriptionEn *string `json:"descriptionEn,omitempty"`
	// Support to copy classes from the source class type. Optional.
	SourceClassType *string `json:"sourceClassType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateClassType instantiates a new CreateClassType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateClassType(name string) *CreateClassType {
	this := CreateClassType{}
	this.Name = name
	return &this
}

// NewCreateClassTypeWithDefaults instantiates a new CreateClassType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateClassTypeWithDefaults() *CreateClassType {
	this := CreateClassType{}
	return &this
}

// GetName returns the Name field value.
func (o *CreateClassType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateClassType) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateClassType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateClassType) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateClassType) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CreateClassType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateClassType) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CreateClassType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayNameEn returns the DisplayNameEn field value if set, zero value otherwise.
func (o *CreateClassType) GetDisplayNameEn() string {
	if o == nil || o.DisplayNameEn == nil {
		var ret string
		return ret
	}
	return *o.DisplayNameEn
}

// GetDisplayNameEnOk returns a tuple with the DisplayNameEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetDisplayNameEnOk() (*string, bool) {
	if o == nil || o.DisplayNameEn == nil {
		return nil, false
	}
	return o.DisplayNameEn, true
}

// HasDisplayNameEn returns a boolean if a field has been set.
func (o *CreateClassType) HasDisplayNameEn() bool {
	return o != nil && o.DisplayNameEn != nil
}

// SetDisplayNameEn gets a reference to the given string and assigns it to the DisplayNameEn field.
func (o *CreateClassType) SetDisplayNameEn(v string) {
	o.DisplayNameEn = &v
}

// GetDescriptionEn returns the DescriptionEn field value if set, zero value otherwise.
func (o *CreateClassType) GetDescriptionEn() string {
	if o == nil || o.DescriptionEn == nil {
		var ret string
		return ret
	}
	return *o.DescriptionEn
}

// GetDescriptionEnOk returns a tuple with the DescriptionEn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetDescriptionEnOk() (*string, bool) {
	if o == nil || o.DescriptionEn == nil {
		return nil, false
	}
	return o.DescriptionEn, true
}

// HasDescriptionEn returns a boolean if a field has been set.
func (o *CreateClassType) HasDescriptionEn() bool {
	return o != nil && o.DescriptionEn != nil
}

// SetDescriptionEn gets a reference to the given string and assigns it to the DescriptionEn field.
func (o *CreateClassType) SetDescriptionEn(v string) {
	o.DescriptionEn = &v
}

// GetSourceClassType returns the SourceClassType field value if set, zero value otherwise.
func (o *CreateClassType) GetSourceClassType() string {
	if o == nil || o.SourceClassType == nil {
		var ret string
		return ret
	}
	return *o.SourceClassType
}

// GetSourceClassTypeOk returns a tuple with the SourceClassType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClassType) GetSourceClassTypeOk() (*string, bool) {
	if o == nil || o.SourceClassType == nil {
		return nil, false
	}
	return o.SourceClassType, true
}

// HasSourceClassType returns a boolean if a field has been set.
func (o *CreateClassType) HasSourceClassType() bool {
	return o != nil && o.SourceClassType != nil
}

// SetSourceClassType gets a reference to the given string and assigns it to the SourceClassType field.
func (o *CreateClassType) SetSourceClassType(v string) {
	o.SourceClassType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateClassType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DisplayNameEn != nil {
		toSerialize["displayNameEn"] = o.DisplayNameEn
	}
	if o.DescriptionEn != nil {
		toSerialize["descriptionEn"] = o.DescriptionEn
	}
	if o.SourceClassType != nil {
		toSerialize["sourceClassType"] = o.SourceClassType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateClassType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string `json:"name"`
		Description     *string `json:"description,omitempty"`
		DisplayName     *string `json:"displayName,omitempty"`
		DisplayNameEn   *string `json:"displayNameEn,omitempty"`
		DescriptionEn   *string `json:"descriptionEn,omitempty"`
		SourceClassType *string `json:"sourceClassType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "displayName", "displayNameEn", "descriptionEn", "sourceClassType"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.DisplayNameEn = all.DisplayNameEn
	o.DescriptionEn = all.DescriptionEn
	o.SourceClassType = all.SourceClassType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
