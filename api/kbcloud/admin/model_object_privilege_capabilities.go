// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeCapabilities struct {
	EngineVersion string                      `json:"engineVersion"`
	Objects       []ObjectPrivilegeCapability `json:"objects"`
	// Native subject kinds supported by this connected engine version.
	SubjectTypes      []ObjectPrivilegeSubjectType `json:"subjectTypes"`
	SubjectQualifiers []ObjectPrivilegeQualifier   `json:"subjectQualifiers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeCapabilities instantiates a new ObjectPrivilegeCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeCapabilities(engineVersion string, objects []ObjectPrivilegeCapability, subjectTypes []ObjectPrivilegeSubjectType) *ObjectPrivilegeCapabilities {
	this := ObjectPrivilegeCapabilities{}
	this.EngineVersion = engineVersion
	this.Objects = objects
	this.SubjectTypes = subjectTypes
	return &this
}

// NewObjectPrivilegeCapabilitiesWithDefaults instantiates a new ObjectPrivilegeCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeCapabilitiesWithDefaults() *ObjectPrivilegeCapabilities {
	this := ObjectPrivilegeCapabilities{}
	return &this
}

// GetEngineVersion returns the EngineVersion field value.
func (o *ObjectPrivilegeCapabilities) GetEngineVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapabilities) GetEngineVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineVersion, true
}

// SetEngineVersion sets field value.
func (o *ObjectPrivilegeCapabilities) SetEngineVersion(v string) {
	o.EngineVersion = v
}

// GetObjects returns the Objects field value.
func (o *ObjectPrivilegeCapabilities) GetObjects() []ObjectPrivilegeCapability {
	if o == nil {
		var ret []ObjectPrivilegeCapability
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapabilities) GetObjectsOk() (*[]ObjectPrivilegeCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Objects, true
}

// SetObjects sets field value.
func (o *ObjectPrivilegeCapabilities) SetObjects(v []ObjectPrivilegeCapability) {
	o.Objects = v
}

// GetSubjectTypes returns the SubjectTypes field value.
func (o *ObjectPrivilegeCapabilities) GetSubjectTypes() []ObjectPrivilegeSubjectType {
	if o == nil {
		var ret []ObjectPrivilegeSubjectType
		return ret
	}
	return o.SubjectTypes
}

// GetSubjectTypesOk returns a tuple with the SubjectTypes field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapabilities) GetSubjectTypesOk() (*[]ObjectPrivilegeSubjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypes, true
}

// SetSubjectTypes sets field value.
func (o *ObjectPrivilegeCapabilities) SetSubjectTypes(v []ObjectPrivilegeSubjectType) {
	o.SubjectTypes = v
}

// GetSubjectQualifiers returns the SubjectQualifiers field value if set, zero value otherwise.
func (o *ObjectPrivilegeCapabilities) GetSubjectQualifiers() []ObjectPrivilegeQualifier {
	if o == nil || o.SubjectQualifiers == nil {
		var ret []ObjectPrivilegeQualifier
		return ret
	}
	return o.SubjectQualifiers
}

// GetSubjectQualifiersOk returns a tuple with the SubjectQualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeCapabilities) GetSubjectQualifiersOk() (*[]ObjectPrivilegeQualifier, bool) {
	if o == nil || o.SubjectQualifiers == nil {
		return nil, false
	}
	return &o.SubjectQualifiers, true
}

// HasSubjectQualifiers returns a boolean if a field has been set.
func (o *ObjectPrivilegeCapabilities) HasSubjectQualifiers() bool {
	return o != nil && o.SubjectQualifiers != nil
}

// SetSubjectQualifiers gets a reference to the given []ObjectPrivilegeQualifier and assigns it to the SubjectQualifiers field.
func (o *ObjectPrivilegeCapabilities) SetSubjectQualifiers(v []ObjectPrivilegeQualifier) {
	o.SubjectQualifiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineVersion"] = o.EngineVersion
	toSerialize["objects"] = o.Objects
	toSerialize["subjectTypes"] = o.SubjectTypes
	if o.SubjectQualifiers != nil {
		toSerialize["subjectQualifiers"] = o.SubjectQualifiers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineVersion     *string                       `json:"engineVersion"`
		Objects           *[]ObjectPrivilegeCapability  `json:"objects"`
		SubjectTypes      *[]ObjectPrivilegeSubjectType `json:"subjectTypes"`
		SubjectQualifiers []ObjectPrivilegeQualifier    `json:"subjectQualifiers,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineVersion == nil {
		return fmt.Errorf("required field engineVersion missing")
	}
	if all.Objects == nil {
		return fmt.Errorf("required field objects missing")
	}
	if all.SubjectTypes == nil {
		return fmt.Errorf("required field subjectTypes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineVersion", "objects", "subjectTypes", "subjectQualifiers"})
	} else {
		return err
	}
	o.EngineVersion = *all.EngineVersion
	o.Objects = *all.Objects
	o.SubjectTypes = *all.SubjectTypes
	o.SubjectQualifiers = all.SubjectQualifiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
