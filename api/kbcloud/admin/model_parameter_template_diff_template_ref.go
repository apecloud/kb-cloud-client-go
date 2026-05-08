// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ParameterTemplateDiffTemplateRef A parameter template reference used by the diff response
type ParameterTemplateDiffTemplateRef struct {
	Name          *string                     `json:"name,omitempty"`
	OrgName       *string                     `json:"orgName,omitempty"`
	Partition     *ParameterTemplatePartition `json:"partition,omitempty"`
	EngineName    *string                     `json:"engineName,omitempty"`
	Component     *string                     `json:"component,omitempty"`
	MajorVersions []string                    `json:"majorVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplateDiffTemplateRef instantiates a new ParameterTemplateDiffTemplateRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplateDiffTemplateRef() *ParameterTemplateDiffTemplateRef {
	this := ParameterTemplateDiffTemplateRef{}
	return &this
}

// NewParameterTemplateDiffTemplateRefWithDefaults instantiates a new ParameterTemplateDiffTemplateRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateDiffTemplateRefWithDefaults() *ParameterTemplateDiffTemplateRef {
	this := ParameterTemplateDiffTemplateRef{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterTemplateDiffTemplateRef) SetName(v string) {
	o.Name = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ParameterTemplateDiffTemplateRef) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetPartition() ParameterTemplatePartition {
	if o == nil || o.Partition == nil {
		var ret ParameterTemplatePartition
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetPartitionOk() (*ParameterTemplatePartition, bool) {
	if o == nil || o.Partition == nil {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasPartition() bool {
	return o != nil && o.Partition != nil
}

// SetPartition gets a reference to the given ParameterTemplatePartition and assigns it to the Partition field.
func (o *ParameterTemplateDiffTemplateRef) SetPartition(v ParameterTemplatePartition) {
	o.Partition = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetEngineName() string {
	if o == nil || o.EngineName == nil {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetEngineNameOk() (*string, bool) {
	if o == nil || o.EngineName == nil {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasEngineName() bool {
	return o != nil && o.EngineName != nil
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *ParameterTemplateDiffTemplateRef) SetEngineName(v string) {
	o.EngineName = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ParameterTemplateDiffTemplateRef) SetComponent(v string) {
	o.Component = &v
}

// GetMajorVersions returns the MajorVersions field value if set, zero value otherwise.
func (o *ParameterTemplateDiffTemplateRef) GetMajorVersions() []string {
	if o == nil || o.MajorVersions == nil {
		var ret []string
		return ret
	}
	return o.MajorVersions
}

// GetMajorVersionsOk returns a tuple with the MajorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiffTemplateRef) GetMajorVersionsOk() (*[]string, bool) {
	if o == nil || o.MajorVersions == nil {
		return nil, false
	}
	return &o.MajorVersions, true
}

// HasMajorVersions returns a boolean if a field has been set.
func (o *ParameterTemplateDiffTemplateRef) HasMajorVersions() bool {
	return o != nil && o.MajorVersions != nil
}

// SetMajorVersions gets a reference to the given []string and assigns it to the MajorVersions field.
func (o *ParameterTemplateDiffTemplateRef) SetMajorVersions(v []string) {
	o.MajorVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplateDiffTemplateRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Partition != nil {
		toSerialize["partition"] = o.Partition
	}
	if o.EngineName != nil {
		toSerialize["engineName"] = o.EngineName
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.MajorVersions != nil {
		toSerialize["majorVersions"] = o.MajorVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplateDiffTemplateRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string                     `json:"name,omitempty"`
		OrgName       *string                     `json:"orgName,omitempty"`
		Partition     *ParameterTemplatePartition `json:"partition,omitempty"`
		EngineName    *string                     `json:"engineName,omitempty"`
		Component     *string                     `json:"component,omitempty"`
		MajorVersions []string                    `json:"majorVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "orgName", "partition", "engineName", "component", "majorVersions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.OrgName = all.OrgName
	if all.Partition != nil && !all.Partition.IsValid() {
		hasInvalidField = true
	} else {
		o.Partition = all.Partition
	}
	o.EngineName = all.EngineName
	o.Component = all.Component
	o.MajorVersions = all.MajorVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
