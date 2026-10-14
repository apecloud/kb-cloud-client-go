// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeSubject struct {
	// Exact native database account or role name.
	Name string `json:"name"`
	// Native authorization subject kind. Does not create accounts or roles, or change memberships.
	Type ObjectPrivilegeSubjectType `json:"type"`
	// Engine-declared native identity qualifiers. MySQL host defaults to %. Unknown keys are rejected, including empty values.
	Qualifiers map[string]string `json:"qualifiers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeSubject instantiates a new ObjectPrivilegeSubject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeSubject(name string, typeVar ObjectPrivilegeSubjectType) *ObjectPrivilegeSubject {
	this := ObjectPrivilegeSubject{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewObjectPrivilegeSubjectWithDefaults instantiates a new ObjectPrivilegeSubject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeSubjectWithDefaults() *ObjectPrivilegeSubject {
	this := ObjectPrivilegeSubject{}
	var typeVar ObjectPrivilegeSubjectType = ObjectPrivilegeSubjectTypeAccount
	this.Type = typeVar
	return &this
}

// GetName returns the Name field value.
func (o *ObjectPrivilegeSubject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSubject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObjectPrivilegeSubject) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *ObjectPrivilegeSubject) GetType() ObjectPrivilegeSubjectType {
	if o == nil {
		var ret ObjectPrivilegeSubjectType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSubject) GetTypeOk() (*ObjectPrivilegeSubjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObjectPrivilegeSubject) SetType(v ObjectPrivilegeSubjectType) {
	o.Type = v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *ObjectPrivilegeSubject) GetQualifiers() map[string]string {
	if o == nil || o.Qualifiers == nil {
		var ret map[string]string
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSubject) GetQualifiersOk() (*map[string]string, bool) {
	if o == nil || o.Qualifiers == nil {
		return nil, false
	}
	return &o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *ObjectPrivilegeSubject) HasQualifiers() bool {
	return o != nil && o.Qualifiers != nil
}

// SetQualifiers gets a reference to the given map[string]string and assigns it to the Qualifiers field.
func (o *ObjectPrivilegeSubject) SetQualifiers(v map[string]string) {
	o.Qualifiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Qualifiers != nil {
		toSerialize["qualifiers"] = o.Qualifiers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeSubject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string                     `json:"name"`
		Type       *ObjectPrivilegeSubjectType `json:"type"`
		Qualifiers map[string]string           `json:"qualifiers,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "qualifiers"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Qualifiers = all.Qualifiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
