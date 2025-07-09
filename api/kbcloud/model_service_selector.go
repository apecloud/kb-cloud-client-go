// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ServiceSelector This definition of component/service/port is in line with kubeblocks api
// appsv1.ServiceRefServiceSelector. Read kubeblocks documentation for more information.
type ServiceSelector struct {
	// referenced cluster's mode
	Mode      string `json:"mode"`
	Component string `json:"component"`
	Service   string `json:"service"`
	Port      string `json:"port"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceSelector instantiates a new ServiceSelector object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceSelector(mode string, component string, service string, port string) *ServiceSelector {
	this := ServiceSelector{}
	this.Mode = mode
	this.Component = component
	this.Service = service
	this.Port = port
	return &this
}

// NewServiceSelectorWithDefaults instantiates a new ServiceSelector object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceSelectorWithDefaults() *ServiceSelector {
	this := ServiceSelector{}
	return &this
}

// GetMode returns the Mode field value.
func (o *ServiceSelector) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ServiceSelector) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ServiceSelector) SetMode(v string) {
	o.Mode = v
}

// GetComponent returns the Component field value.
func (o *ServiceSelector) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ServiceSelector) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ServiceSelector) SetComponent(v string) {
	o.Component = v
}

// GetService returns the Service field value.
func (o *ServiceSelector) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ServiceSelector) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ServiceSelector) SetService(v string) {
	o.Service = v
}

// GetPort returns the Port field value.
func (o *ServiceSelector) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ServiceSelector) GetPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value.
func (o *ServiceSelector) SetPort(v string) {
	o.Port = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["component"] = o.Component
	toSerialize["service"] = o.Service
	toSerialize["port"] = o.Port

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceSelector) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode      *string `json:"mode"`
		Component *string `json:"component"`
		Service   *string `json:"service"`
		Port      *string `json:"port"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.Port == nil {
		return fmt.Errorf("required field port missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "component", "service", "port"})
	} else {
		return err
	}
	o.Mode = *all.Mode
	o.Component = *all.Component
	o.Service = *all.Service
	o.Port = *all.Port

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
