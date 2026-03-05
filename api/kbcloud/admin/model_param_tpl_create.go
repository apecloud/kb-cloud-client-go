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
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// Name of the original parameter template.
	OriName      string                     `json:"oriName"`
	OriPartition ParameterTemplatePartition `json:"oriPartition"`
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
func NewParamTplCreate(description string, name string, oriName string, oriPartition ParameterTemplatePartition) *ParamTplCreate {
	this := ParamTplCreate{}
	this.Description = description
	this.Name = name
	this.OriName = oriName
	this.OriPartition = oriPartition
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

// GetOriName returns the OriName field value.
func (o *ParamTplCreate) GetOriName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OriName
}

// GetOriNameOk returns a tuple with the OriName field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetOriNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriName, true
}

// SetOriName sets field value.
func (o *ParamTplCreate) SetOriName(v string) {
	o.OriName = v
}

// GetOriPartition returns the OriPartition field value.
func (o *ParamTplCreate) GetOriPartition() ParameterTemplatePartition {
	if o == nil {
		var ret ParameterTemplatePartition
		return ret
	}
	return o.OriPartition
}

// GetOriPartitionOk returns a tuple with the OriPartition field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetOriPartitionOk() (*ParameterTemplatePartition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriPartition, true
}

// SetOriPartition sets field value.
func (o *ParamTplCreate) SetOriPartition(v ParameterTemplatePartition) {
	o.OriPartition = v
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
	toSerialize["name"] = o.Name
	toSerialize["oriName"] = o.OriName
	toSerialize["oriPartition"] = o.OriPartition
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
		Description  *string                     `json:"description"`
		Name         *string                     `json:"name"`
		OriName      *string                     `json:"oriName"`
		OriPartition *ParameterTemplatePartition `json:"oriPartition"`
		IsPrivate    *bool                       `json:"isPrivate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OriName == nil {
		return fmt.Errorf("required field oriName missing")
	}
	if all.OriPartition == nil {
		return fmt.Errorf("required field oriPartition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "name", "oriName", "oriPartition", "isPrivate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Name = *all.Name
	o.OriName = *all.OriName
	if !all.OriPartition.IsValid() {
		hasInvalidField = true
	} else {
		o.OriPartition = *all.OriPartition
	}
	o.IsPrivate = all.IsPrivate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
