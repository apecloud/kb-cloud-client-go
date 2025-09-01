// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplCreate parameter template create
type ParamTplCreate struct {
	// Description of parameter template
	Description string `json:"description"`
	// Major version of database engine, eg: 8.0
	MajorVersion string `json:"majorVersion"`
	// Name of database engine
	Engine string `json:"engine"`
	// Name of component
	Component string `json:"component"`
	// Name of parameter template. Name must be unique within an Org
	Name string `json:"name"`
	// Name of custom parameter template. When set customName, will create a copy of this custom parameter template.
	CustomName   *string            `json:"customName,omitempty"`
	OriPartition *ParamTplPartition `json:"oriPartition,omitempty"`
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
func NewParamTplCreate(description string, majorVersion string, engine string, component string, name string) *ParamTplCreate {
	this := ParamTplCreate{}
	this.Description = description
	this.MajorVersion = majorVersion
	this.Engine = engine
	this.Component = component
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

// GetMajorVersion returns the MajorVersion field value.
func (o *ParamTplCreate) GetMajorVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetMajorVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersion, true
}

// SetMajorVersion sets field value.
func (o *ParamTplCreate) SetMajorVersion(v string) {
	o.MajorVersion = v
}

// GetEngine returns the Engine field value.
func (o *ParamTplCreate) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ParamTplCreate) SetEngine(v string) {
	o.Engine = v
}

// GetComponent returns the Component field value.
func (o *ParamTplCreate) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParamTplCreate) SetComponent(v string) {
	o.Component = v
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

// GetOriPartition returns the OriPartition field value if set, zero value otherwise.
func (o *ParamTplCreate) GetOriPartition() ParamTplPartition {
	if o == nil || o.OriPartition == nil {
		var ret ParamTplPartition
		return ret
	}
	return *o.OriPartition
}

// GetOriPartitionOk returns a tuple with the OriPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplCreate) GetOriPartitionOk() (*ParamTplPartition, bool) {
	if o == nil || o.OriPartition == nil {
		return nil, false
	}
	return o.OriPartition, true
}

// HasOriPartition returns a boolean if a field has been set.
func (o *ParamTplCreate) HasOriPartition() bool {
	return o != nil && o.OriPartition != nil
}

// SetOriPartition gets a reference to the given ParamTplPartition and assigns it to the OriPartition field.
func (o *ParamTplCreate) SetOriPartition(v ParamTplPartition) {
	o.OriPartition = &v
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
	toSerialize["majorVersion"] = o.MajorVersion
	toSerialize["engine"] = o.Engine
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name
	if o.CustomName != nil {
		toSerialize["customName"] = o.CustomName
	}
	if o.OriPartition != nil {
		toSerialize["oriPartition"] = o.OriPartition
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
		Description  *string            `json:"description"`
		MajorVersion *string            `json:"majorVersion"`
		Engine       *string            `json:"engine"`
		Component    *string            `json:"component"`
		Name         *string            `json:"name"`
		CustomName   *string            `json:"customName,omitempty"`
		OriPartition *ParamTplPartition `json:"oriPartition,omitempty"`
		IsPrivate    *bool              `json:"isPrivate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.MajorVersion == nil {
		return fmt.Errorf("required field majorVersion missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "majorVersion", "engine", "component", "name", "customName", "oriPartition", "isPrivate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.MajorVersion = *all.MajorVersion
	o.Engine = *all.Engine
	o.Component = *all.Component
	o.Name = *all.Name
	o.CustomName = all.CustomName
	if all.OriPartition != nil && !all.OriPartition.IsValid() {
		hasInvalidField = true
	} else {
		o.OriPartition = all.OriPartition
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
