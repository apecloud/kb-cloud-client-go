// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplCreateFromCluster parameter template create

type ParamTplCreateFromCluster struct {
	// Description of parameter template
	Description string `json:"description"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// component type.
	Component *string `json:"component,omitempty"`
	// database engine Version
	EngineVersion *string `json:"engineVersion,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplCreateFromCluster instantiates a new ParamTplCreateFromCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplCreateFromCluster(description string, name string) *ParamTplCreateFromCluster {
	this := ParamTplCreateFromCluster{}
	this.Description = description
	this.Name = name
	return &this
}

// NewParamTplCreateFromClusterWithDefaults instantiates a new ParamTplCreateFromCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplCreateFromClusterWithDefaults() *ParamTplCreateFromCluster {
	this := ParamTplCreateFromCluster{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ParamTplCreateFromCluster) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreateFromCluster) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ParamTplCreateFromCluster) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *ParamTplCreateFromCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreateFromCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParamTplCreateFromCluster) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ParamTplCreateFromCluster) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplCreateFromCluster) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ParamTplCreateFromCluster) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ParamTplCreateFromCluster) SetComponent(v string) {
	o.Component = &v
}

// GetEngineVersion returns the EngineVersion field value if set, zero value otherwise.
func (o *ParamTplCreateFromCluster) GetEngineVersion() string {
	if o == nil || o.EngineVersion == nil {
		var ret string
		return ret
	}
	return *o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplCreateFromCluster) GetEngineVersionOk() (*string, bool) {
	if o == nil || o.EngineVersion == nil {
		return nil, false
	}
	return o.EngineVersion, true
}

// HasEngineVersion returns a boolean if a field has been set.
func (o *ParamTplCreateFromCluster) HasEngineVersion() bool {
	return o != nil && o.EngineVersion != nil
}

// SetEngineVersion gets a reference to the given string and assigns it to the EngineVersion field.
func (o *ParamTplCreateFromCluster) SetEngineVersion(v string) {
	o.EngineVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplCreateFromCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.EngineVersion != nil {
		toSerialize["engineVersion"] = o.EngineVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplCreateFromCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string `json:"description"`
		Name          *string `json:"name"`
		Component     *string `json:"component,omitempty"`
		EngineVersion *string `json:"engineVersion,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "name", "component", "engineVersion"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Name = *all.Name
	o.Component = all.Component
	o.EngineVersion = all.EngineVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
