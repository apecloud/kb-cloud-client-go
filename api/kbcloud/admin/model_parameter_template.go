// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterTemplate struct {
	// name of default parameter template
	Name        string               `json:"name"`
	Description LocalizedDescription `json:"description"`
	// match the major version set in the component
	MajorVersion *string `json:"majorVersion,omitempty"`
	// match the major versions set in the component, such as 8.0, 8.1, 8.2
	MajorVersions []string `json:"majorVersions,omitempty"`
	// refs contains the references to the configuration templates
	Refs []Ref `json:"refs"`
	// whether the default parameter template is used by default, set in componentDefinition.configs, only one default parameter template can be set in certain family.
	DefaultUse bool `json:"defaultUse"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplate instantiates a new ParameterTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplate(name string, description LocalizedDescription, refs []Ref, defaultUse bool) *ParameterTemplate {
	this := ParameterTemplate{}
	this.Name = name
	this.Description = description
	this.Refs = refs
	this.DefaultUse = defaultUse
	return &this
}

// NewParameterTemplateWithDefaults instantiates a new ParameterTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateWithDefaults() *ParameterTemplate {
	this := ParameterTemplate{}
	return &this
}

// GetName returns the Name field value.
func (o *ParameterTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParameterTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *ParameterTemplate) GetDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParameterTemplate) SetDescription(v LocalizedDescription) {
	o.Description = v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *ParameterTemplate) GetMajorVersion() string {
	if o == nil || o.MajorVersion == nil {
		var ret string
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetMajorVersionOk() (*string, bool) {
	if o == nil || o.MajorVersion == nil {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *ParameterTemplate) HasMajorVersion() bool {
	return o != nil && o.MajorVersion != nil
}

// SetMajorVersion gets a reference to the given string and assigns it to the MajorVersion field.
func (o *ParameterTemplate) SetMajorVersion(v string) {
	o.MajorVersion = &v
}

// GetMajorVersions returns the MajorVersions field value if set, zero value otherwise.
func (o *ParameterTemplate) GetMajorVersions() []string {
	if o == nil || o.MajorVersions == nil {
		var ret []string
		return ret
	}
	return o.MajorVersions
}

// GetMajorVersionsOk returns a tuple with the MajorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetMajorVersionsOk() (*[]string, bool) {
	if o == nil || o.MajorVersions == nil {
		return nil, false
	}
	return &o.MajorVersions, true
}

// HasMajorVersions returns a boolean if a field has been set.
func (o *ParameterTemplate) HasMajorVersions() bool {
	return o != nil && o.MajorVersions != nil
}

// SetMajorVersions gets a reference to the given []string and assigns it to the MajorVersions field.
func (o *ParameterTemplate) SetMajorVersions(v []string) {
	o.MajorVersions = v
}

// GetRefs returns the Refs field value.
func (o *ParameterTemplate) GetRefs() []Ref {
	if o == nil {
		var ret []Ref
		return ret
	}
	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetRefsOk() (*[]Ref, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Refs, true
}

// SetRefs sets field value.
func (o *ParameterTemplate) SetRefs(v []Ref) {
	o.Refs = v
}

// GetDefaultUse returns the DefaultUse field value.
func (o *ParameterTemplate) GetDefaultUse() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DefaultUse
}

// GetDefaultUseOk returns a tuple with the DefaultUse field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplate) GetDefaultUseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultUse, true
}

// SetDefaultUse sets field value.
func (o *ParameterTemplate) SetDefaultUse(v bool) {
	o.DefaultUse = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if o.MajorVersion != nil {
		toSerialize["majorVersion"] = o.MajorVersion
	}
	if o.MajorVersions != nil {
		toSerialize["majorVersions"] = o.MajorVersions
	}
	toSerialize["refs"] = o.Refs
	toSerialize["defaultUse"] = o.DefaultUse

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string               `json:"name"`
		Description   *LocalizedDescription `json:"description"`
		MajorVersion  *string               `json:"majorVersion,omitempty"`
		MajorVersions []string              `json:"majorVersions,omitempty"`
		Refs          *[]Ref                `json:"refs"`
		DefaultUse    *bool                 `json:"defaultUse"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Refs == nil {
		return fmt.Errorf("required field refs missing")
	}
	if all.DefaultUse == nil {
		return fmt.Errorf("required field defaultUse missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "majorVersion", "majorVersions", "refs", "defaultUse"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = *all.Description
	o.MajorVersion = all.MajorVersion
	o.MajorVersions = all.MajorVersions
	o.Refs = *all.Refs
	o.DefaultUse = *all.DefaultUse

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
