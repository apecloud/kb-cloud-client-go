// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InstanceTemplateUpgrade struct {
	Name           string `json:"name"`
	ServiceVersion string `json:"serviceVersion"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceTemplateUpgrade instantiates a new InstanceTemplateUpgrade object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceTemplateUpgrade(name string, serviceVersion string) *InstanceTemplateUpgrade {
	this := InstanceTemplateUpgrade{}
	this.Name = name
	this.ServiceVersion = serviceVersion
	return &this
}

// NewInstanceTemplateUpgradeWithDefaults instantiates a new InstanceTemplateUpgrade object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceTemplateUpgradeWithDefaults() *InstanceTemplateUpgrade {
	this := InstanceTemplateUpgrade{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceTemplateUpgrade) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateUpgrade) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceTemplateUpgrade) SetName(v string) {
	o.Name = v
}

// GetServiceVersion returns the ServiceVersion field value.
func (o *InstanceTemplateUpgrade) GetServiceVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value
// and a boolean to check if the value has been set.
func (o *InstanceTemplateUpgrade) GetServiceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceVersion, true
}

// SetServiceVersion sets field value.
func (o *InstanceTemplateUpgrade) SetServiceVersion(v string) {
	o.ServiceVersion = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceTemplateUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["serviceVersion"] = o.ServiceVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceTemplateUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string `json:"name"`
		ServiceVersion *string `json:"serviceVersion"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ServiceVersion == nil {
		return fmt.Errorf("required field serviceVersion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "serviceVersion"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.ServiceVersion = *all.ServiceVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
