// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineTypeItem struct {
	// engine type name
	Name        string               `json:"name"`
	DisplayName LocalizedDescription `json:"displayName"`
	Description LocalizedDescription `json:"description"`
	// engine type display order
	DisplayOrder int32 `json:"displayOrder"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineTypeItem instantiates a new EngineTypeItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineTypeItem(name string, displayName LocalizedDescription, description LocalizedDescription, displayOrder int32) *EngineTypeItem {
	this := EngineTypeItem{}
	this.Name = name
	this.DisplayName = displayName
	this.Description = description
	this.DisplayOrder = displayOrder
	return &this
}

// NewEngineTypeItemWithDefaults instantiates a new EngineTypeItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineTypeItemWithDefaults() *EngineTypeItem {
	this := EngineTypeItem{}
	return &this
}

// GetName returns the Name field value.
func (o *EngineTypeItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineTypeItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineTypeItem) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value.
func (o *EngineTypeItem) GetDisplayName() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EngineTypeItem) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *EngineTypeItem) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = v
}

// GetDescription returns the Description field value.
func (o *EngineTypeItem) GetDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EngineTypeItem) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *EngineTypeItem) SetDescription(v LocalizedDescription) {
	o.Description = v
}

// GetDisplayOrder returns the DisplayOrder field value.
func (o *EngineTypeItem) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *EngineTypeItem) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value.
func (o *EngineTypeItem) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineTypeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["displayOrder"] = o.DisplayOrder

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineTypeItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string               `json:"name"`
		DisplayName  *LocalizedDescription `json:"displayName"`
		Description  *LocalizedDescription `json:"description"`
		DisplayOrder *int32                `json:"displayOrder"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayOrder == nil {
		return fmt.Errorf("required field displayOrder missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "displayName", "description", "displayOrder"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = *all.DisplayName
	if all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = *all.Description
	o.DisplayOrder = *all.DisplayOrder

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
