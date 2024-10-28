// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformParameter platformParameter item
type PlatformParameter struct {
	// platformParameter ID
	Id string `json:"id"`
	// platformParameter name
	Name string `json:"name"`
	// platformParameter category, representing the platformParameter group, such as 'recycle bin'
	Category PlatformParameterCategory `json:"category"`
	// type of platformParameter value
	Type string `json:"type"`
	// platformParameter constraints including min, max, enum, default value
	Constraints PlatformParameterConstraints `json:"constraints"`
	// platformParameter value
	Value *string `json:"value,omitempty"`
	// platformParameter description
	Description string `json:"description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformParameter instantiates a new PlatformParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformParameter(id string, name string, category PlatformParameterCategory, typeVar string, constraints PlatformParameterConstraints, description string) *PlatformParameter {
	this := PlatformParameter{}
	this.Id = id
	this.Name = name
	this.Category = category
	this.Type = typeVar
	this.Constraints = constraints
	this.Description = description
	return &this
}

// NewPlatformParameterWithDefaults instantiates a new PlatformParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformParameterWithDefaults() *PlatformParameter {
	this := PlatformParameter{}
	return &this
}

// GetId returns the Id field value.
func (o *PlatformParameter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PlatformParameter) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *PlatformParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *PlatformParameter) SetName(v string) {
	o.Name = v
}

// GetCategory returns the Category field value.
func (o *PlatformParameter) GetCategory() PlatformParameterCategory {
	if o == nil {
		var ret PlatformParameterCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetCategoryOk() (*PlatformParameterCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *PlatformParameter) SetCategory(v PlatformParameterCategory) {
	o.Category = v
}

// GetType returns the Type field value.
func (o *PlatformParameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PlatformParameter) SetType(v string) {
	o.Type = v
}

// GetConstraints returns the Constraints field value.
func (o *PlatformParameter) GetConstraints() PlatformParameterConstraints {
	if o == nil {
		var ret PlatformParameterConstraints
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetConstraintsOk() (*PlatformParameterConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraints, true
}

// SetConstraints sets field value.
func (o *PlatformParameter) SetConstraints(v PlatformParameterConstraints) {
	o.Constraints = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PlatformParameter) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlatformParameter) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PlatformParameter) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value.
func (o *PlatformParameter) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PlatformParameter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *PlatformParameter) SetDescription(v string) {
	o.Description = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["category"] = o.Category
	toSerialize["type"] = o.Type
	toSerialize["constraints"] = o.Constraints
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string                       `json:"id"`
		Name        *string                       `json:"name"`
		Category    *PlatformParameterCategory    `json:"category"`
		Type        *string                       `json:"type"`
		Constraints *PlatformParameterConstraints `json:"constraints"`
		Value       *string                       `json:"value,omitempty"`
		Description *string                       `json:"description"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "category", "type", "constraints", "value", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = *all.Name
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Type = *all.Type
	if all.Constraints.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Constraints = *all.Constraints
	o.Value = all.Value
	o.Description = *all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
