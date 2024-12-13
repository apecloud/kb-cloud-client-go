// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// OrgParameter org parameter item
type OrgParameter struct {
	// org parameter id
	Id string `json:"id"`
	// org id
	OrgId string `json:"orgId"`
	// org parameter name
	Name string `json:"name"`
	// org parameter category
	Category string `json:"category"`
	// type of org parameter value
	Type string `json:"type"`
	// org parameter constraints including min, max, enum, default value
	Constraints OrgParameterConstraints `json:"constraints"`
	// org parameter description
	Description string `json:"description"`
	// org parameter value
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgParameter instantiates a new OrgParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgParameter(id string, orgId string, name string, category string, typeVar string, constraints OrgParameterConstraints, description string, value string) *OrgParameter {
	this := OrgParameter{}
	this.Id = id
	this.OrgId = orgId
	this.Name = name
	this.Category = category
	this.Type = typeVar
	this.Constraints = constraints
	this.Description = description
	this.Value = value
	return &this
}

// NewOrgParameterWithDefaults instantiates a new OrgParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgParameterWithDefaults() *OrgParameter {
	this := OrgParameter{}
	return &this
}

// GetId returns the Id field value.
func (o *OrgParameter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OrgParameter) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value.
func (o *OrgParameter) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *OrgParameter) SetOrgId(v string) {
	o.OrgId = v
}

// GetName returns the Name field value.
func (o *OrgParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OrgParameter) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value.
func (o *OrgParameter) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *OrgParameter) SetCategory(v string) {
	o.Category = v
}

// GetType returns the Type field value.
func (o *OrgParameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgParameter) SetType(v string) {
	o.Type = v
}

// GetConstraints returns the Constraints field value.
func (o *OrgParameter) GetConstraints() OrgParameterConstraints {
	if o == nil {
		var ret OrgParameterConstraints
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetConstraintsOk() (*OrgParameterConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraints, true
}

// SetConstraints sets field value.
func (o *OrgParameter) SetConstraints(v OrgParameterConstraints) {
	o.Constraints = v
}

// GetDescription returns the Description field value.
func (o *OrgParameter) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *OrgParameter) SetDescription(v string) {
	o.Description = v
}

// GetValue returns the Value field value.
func (o *OrgParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OrgParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *OrgParameter) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["orgId"] = o.OrgId
	toSerialize["name"] = o.Name
	toSerialize["category"] = o.Category
	toSerialize["type"] = o.Type
	toSerialize["constraints"] = o.Constraints
	toSerialize["description"] = o.Description
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string                  `json:"id"`
		OrgId       *string                  `json:"orgId"`
		Name        *string                  `json:"name"`
		Category    *string                  `json:"category"`
		Type        *string                  `json:"type"`
		Constraints *OrgParameterConstraints `json:"constraints"`
		Description *string                  `json:"description"`
		Value       *string                  `json:"value"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field orgId missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Constraints == nil {
		return fmt.Errorf("required field constraints missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgId", "name", "category", "type", "constraints", "description", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.OrgId = *all.OrgId
	o.Name = *all.Name
	o.Category = *all.Category
	o.Type = *all.Type
	if all.Constraints.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Constraints = *all.Constraints
	o.Description = *all.Description
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
