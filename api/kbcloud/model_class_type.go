// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClassType struct {
	// The unique identifier for the class type.
	Id *string `json:"id,omitempty"`
	// Consists of lowercase English letters and numbers, 1-8 characters long. Used to add to the class type name.
	Name        *string               `json:"name,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	DisplayName *LocalizedDescription `json:"displayName,omitempty"`
	// Creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClassType instantiates a new ClassType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClassType() *ClassType {
	this := ClassType{}
	return &this
}

// NewClassTypeWithDefaults instantiates a new ClassType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClassTypeWithDefaults() *ClassType {
	this := ClassType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClassType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClassType) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClassType) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClassType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClassType) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClassType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClassType) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassType) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClassType) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *ClassType) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClassType) GetDisplayName() LocalizedDescription {
	if o == nil || o.DisplayName == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassType) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClassType) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given LocalizedDescription and assigns it to the DisplayName field.
func (o *ClassType) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ClassType) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ClassType) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ClassType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClassType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClassType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string               `json:"id,omitempty"`
		Name        *string               `json:"name,omitempty"`
		Description *LocalizedDescription `json:"description,omitempty"`
		DisplayName *LocalizedDescription `json:"displayName,omitempty"`
		CreatedAt   *time.Time            `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "description", "displayName", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	if all.DisplayName != nil && all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = all.DisplayName
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
