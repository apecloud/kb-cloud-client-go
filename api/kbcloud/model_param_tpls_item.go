// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplsItem the item of the parameter template
type ParamTplsItem struct {
	// component type, refer to componentDef and support NamePrefix
	Component *string `json:"component,omitempty"`
	// Name of the master parameter template to create a cluster-specific template from.
	// This field always holds the name of the master template.
	ParamTplName *string `json:"paramTplName,omitempty"`
	// Name of the parameter template instance actually applied to the cluster. This is a system-managed, read-only field.
	AppliedParamTplName *string `json:"appliedParamTplName,omitempty"`
	// id of parameter template
	ParamTplId common.NullableString `json:"paramTplID,omitempty"`
	// org name of the parameter template
	OrgName           common.NullableString       `json:"orgName,omitempty"`
	ParamTplPartition *ParameterTemplatePartition `json:"paramTplPartition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplsItem instantiates a new ParamTplsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplsItem() *ParamTplsItem {
	this := ParamTplsItem{}
	return &this
}

// NewParamTplsItemWithDefaults instantiates a new ParamTplsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplsItemWithDefaults() *ParamTplsItem {
	this := ParamTplsItem{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParamTplsItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParamTplsItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ParamTplsItem) SetComponent(v string) {
	o.Component = &v
}

// GetParamTplName returns the ParamTplName field value if set, zero value otherwise.
func (o *ParamTplsItem) GetParamTplName() string {
	if o == nil || o.ParamTplName == nil {
		var ret string
		return ret
	}
	return *o.ParamTplName
}

// GetParamTplNameOk returns a tuple with the ParamTplName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetParamTplNameOk() (*string, bool) {
	if o == nil || o.ParamTplName == nil {
		return nil, false
	}
	return o.ParamTplName, true
}

// HasParamTplName returns a boolean if a field has been set.
func (o *ParamTplsItem) HasParamTplName() bool {
	return o != nil && o.ParamTplName != nil
}

// SetParamTplName gets a reference to the given string and assigns it to the ParamTplName field.
func (o *ParamTplsItem) SetParamTplName(v string) {
	o.ParamTplName = &v
}

// GetAppliedParamTplName returns the AppliedParamTplName field value if set, zero value otherwise.
func (o *ParamTplsItem) GetAppliedParamTplName() string {
	if o == nil || o.AppliedParamTplName == nil {
		var ret string
		return ret
	}
	return *o.AppliedParamTplName
}

// GetAppliedParamTplNameOk returns a tuple with the AppliedParamTplName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetAppliedParamTplNameOk() (*string, bool) {
	if o == nil || o.AppliedParamTplName == nil {
		return nil, false
	}
	return o.AppliedParamTplName, true
}

// HasAppliedParamTplName returns a boolean if a field has been set.
func (o *ParamTplsItem) HasAppliedParamTplName() bool {
	return o != nil && o.AppliedParamTplName != nil
}

// SetAppliedParamTplName gets a reference to the given string and assigns it to the AppliedParamTplName field.
func (o *ParamTplsItem) SetAppliedParamTplName(v string) {
	o.AppliedParamTplName = &v
}

// GetParamTplId returns the ParamTplId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParamTplsItem) GetParamTplId() string {
	if o == nil || o.ParamTplId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParamTplId.Get()
}

// GetParamTplIdOk returns a tuple with the ParamTplId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ParamTplsItem) GetParamTplIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParamTplId.Get(), o.ParamTplId.IsSet()
}

// HasParamTplId returns a boolean if a field has been set.
func (o *ParamTplsItem) HasParamTplId() bool {
	return o != nil && o.ParamTplId.IsSet()
}

// SetParamTplId gets a reference to the given common.NullableString and assigns it to the ParamTplId field.
func (o *ParamTplsItem) SetParamTplId(v string) {
	o.ParamTplId.Set(&v)
}

// SetParamTplIdNil sets the value for ParamTplId to be an explicit nil.
func (o *ParamTplsItem) SetParamTplIdNil() {
	o.ParamTplId.Set(nil)
}

// UnsetParamTplId ensures that no value is present for ParamTplId, not even an explicit nil.
func (o *ParamTplsItem) UnsetParamTplId() {
	o.ParamTplId.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParamTplsItem) GetOrgName() string {
	if o == nil || o.OrgName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgName.Get()
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ParamTplsItem) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgName.Get(), o.OrgName.IsSet()
}

// HasOrgName returns a boolean if a field has been set.
func (o *ParamTplsItem) HasOrgName() bool {
	return o != nil && o.OrgName.IsSet()
}

// SetOrgName gets a reference to the given common.NullableString and assigns it to the OrgName field.
func (o *ParamTplsItem) SetOrgName(v string) {
	o.OrgName.Set(&v)
}

// SetOrgNameNil sets the value for OrgName to be an explicit nil.
func (o *ParamTplsItem) SetOrgNameNil() {
	o.OrgName.Set(nil)
}

// UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil.
func (o *ParamTplsItem) UnsetOrgName() {
	o.OrgName.Unset()
}

// GetParamTplPartition returns the ParamTplPartition field value if set, zero value otherwise.
func (o *ParamTplsItem) GetParamTplPartition() ParameterTemplatePartition {
	if o == nil || o.ParamTplPartition == nil {
		var ret ParameterTemplatePartition
		return ret
	}
	return *o.ParamTplPartition
}

// GetParamTplPartitionOk returns a tuple with the ParamTplPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplsItem) GetParamTplPartitionOk() (*ParameterTemplatePartition, bool) {
	if o == nil || o.ParamTplPartition == nil {
		return nil, false
	}
	return o.ParamTplPartition, true
}

// HasParamTplPartition returns a boolean if a field has been set.
func (o *ParamTplsItem) HasParamTplPartition() bool {
	return o != nil && o.ParamTplPartition != nil
}

// SetParamTplPartition gets a reference to the given ParameterTemplatePartition and assigns it to the ParamTplPartition field.
func (o *ParamTplsItem) SetParamTplPartition(v ParameterTemplatePartition) {
	o.ParamTplPartition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ParamTplName != nil {
		toSerialize["paramTplName"] = o.ParamTplName
	}
	if o.AppliedParamTplName != nil {
		toSerialize["appliedParamTplName"] = o.AppliedParamTplName
	}
	if o.ParamTplId.IsSet() {
		toSerialize["paramTplID"] = o.ParamTplId.Get()
	}
	if o.OrgName.IsSet() {
		toSerialize["orgName"] = o.OrgName.Get()
	}
	if o.ParamTplPartition != nil {
		toSerialize["paramTplPartition"] = o.ParamTplPartition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component           *string                     `json:"component,omitempty"`
		ParamTplName        *string                     `json:"paramTplName,omitempty"`
		AppliedParamTplName *string                     `json:"appliedParamTplName,omitempty"`
		ParamTplId          common.NullableString       `json:"paramTplID,omitempty"`
		OrgName             common.NullableString       `json:"orgName,omitempty"`
		ParamTplPartition   *ParameterTemplatePartition `json:"paramTplPartition,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "paramTplName", "appliedParamTplName", "paramTplID", "orgName", "paramTplPartition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = all.Component
	o.ParamTplName = all.ParamTplName
	o.AppliedParamTplName = all.AppliedParamTplName
	o.ParamTplId = all.ParamTplId
	o.OrgName = all.OrgName
	if all.ParamTplPartition != nil && !all.ParamTplPartition.IsValid() {
		hasInvalidField = true
	} else {
		o.ParamTplPartition = all.ParamTplPartition
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
