// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplCreate parameter template create
type ParamTplCreate struct {
	// Description of parameter template
	Description string `json:"description"`
	// Name of database with its version, eg: mysql8.0
	Family string `json:"family"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// Name of custom parameter template. When set customName, will create a copy of this custom parameter template.
	CustomName *string `json:"customName,omitempty"`
	// Determines whether the user can see this parameter template
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplCreate instantiates a new ParamTplCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplCreate(description string, family string, name string) *ParamTplCreate {
	this := ParamTplCreate{}
	this.Description = description
	this.Family = family
	this.Name = name
	return &this
}

// NewParamTplCreateWithDefaults instantiates a new ParamTplCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplCreateWithDefaults() *ParamTplCreate {
	this := ParamTplCreate{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ParamTplCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParamTplCreate) SetDescription(v string) {
	o.Description = v
}

// GetFamily returns the Family field value.
func (o *ParamTplCreate) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value.
func (o *ParamTplCreate) SetFamily(v string) {
	o.Family = v
}

// GetName returns the Name field value.
func (o *ParamTplCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParamTplCreate) SetName(v string) {
	o.Name = v
}

// GetCustomName returns the CustomName field value if set, zero value otherwise.
func (o *ParamTplCreate) GetCustomName() string {
	if o == nil || o.CustomName == nil {
		var ret string
		return ret
	}
	return *o.CustomName
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetCustomNameOk() (*string, bool) {
	if o == nil || o.CustomName == nil {
		return nil, false
	}
	return o.CustomName, true
}

// HasCustomName returns a boolean if a field has been set.
func (o *ParamTplCreate) HasCustomName() bool {
	return o != nil && o.CustomName != nil
}

// SetCustomName gets a reference to the given string and assigns it to the CustomName field.
func (o *ParamTplCreate) SetCustomName(v string) {
	o.CustomName = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *ParamTplCreate) GetIsPrivate() bool {
	if o == nil || o.IsPrivate == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetIsPrivateOk() (*bool, bool) {
	if o == nil || o.IsPrivate == nil {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *ParamTplCreate) HasIsPrivate() bool {
	return o != nil && o.IsPrivate != nil
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *ParamTplCreate) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["family"] = o.Family
	toSerialize["name"] = o.Name
	if o.CustomName != nil {
		toSerialize["customName"] = o.CustomName
	}
	if o.IsPrivate != nil {
		toSerialize["isPrivate"] = o.IsPrivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description"`
		Family      *string `json:"family"`
		Name        *string `json:"name"`
		CustomName  *string `json:"customName,omitempty"`
		IsPrivate   *bool   `json:"isPrivate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Family == nil {
		return fmt.Errorf("required field family missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "family", "name", "customName", "isPrivate"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Family = *all.Family
	o.Name = *all.Name
	o.CustomName = all.CustomName
	o.IsPrivate = all.IsPrivate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
