// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Endpoint Endpoint is the information of cluster endpoints

type Endpoint struct {
	// title of the endpoint
	Title string `json:"title"`
	// Component type name
	Component string `json:"component"`
	// Endpoint hosts
	Hosts []string `json:"hosts"`
	// Endpoint port
	Port int32 `json:"port"`
	// Type of endpoint
	Type EndpointType `json:"type"`
	// Network type of endpoint
	NetworkType EndpointNetworkType `json:"networkType"`
	// Service name of endpoint
	ServiceName string `json:"serviceName"`
	// Port name of endpoint
	PortName string `json:"portName"`
	// Endpoint backend instances
	Instances []string `json:"instances,omitempty"`
	// Whether the endpoint is mutable
	Mutable bool `json:"mutable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEndpoint instantiates a new Endpoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEndpoint(title string, component string, hosts []string, port int32, typeVar EndpointType, networkType EndpointNetworkType, serviceName string, portName string, mutable bool) *Endpoint {
	this := Endpoint{}
	this.Title = title
	this.Component = component
	this.Hosts = hosts
	this.Port = port
	this.Type = typeVar
	this.NetworkType = networkType
	this.ServiceName = serviceName
	this.PortName = portName
	this.Mutable = mutable
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetTitle returns the Title field value.
func (o *Endpoint) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *Endpoint) SetTitle(v string) {
	o.Title = v
}

// GetComponent returns the Component field value.
func (o *Endpoint) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *Endpoint) SetComponent(v string) {
	o.Component = v
}

// GetHosts returns the Hosts field value.
func (o *Endpoint) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetHostsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// SetHosts sets field value.
func (o *Endpoint) SetHosts(v []string) {
	o.Hosts = v
}

// GetPort returns the Port field value.
func (o *Endpoint) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value.
func (o *Endpoint) SetPort(v int32) {
	o.Port = v
}

// GetType returns the Type field value.
func (o *Endpoint) GetType() EndpointType {
	if o == nil {
		var ret EndpointType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetTypeOk() (*EndpointType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Endpoint) SetType(v EndpointType) {
	o.Type = v
}

// GetNetworkType returns the NetworkType field value.
func (o *Endpoint) GetNetworkType() EndpointNetworkType {
	if o == nil {
		var ret EndpointNetworkType
		return ret
	}
	return o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetNetworkTypeOk() (*EndpointNetworkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkType, true
}

// SetNetworkType sets field value.
func (o *Endpoint) SetNetworkType(v EndpointNetworkType) {
	o.NetworkType = v
}

// GetServiceName returns the ServiceName field value.
func (o *Endpoint) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value.
func (o *Endpoint) SetServiceName(v string) {
	o.ServiceName = v
}

// GetPortName returns the PortName field value.
func (o *Endpoint) GetPortName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortName, true
}

// SetPortName sets field value.
func (o *Endpoint) SetPortName(v string) {
	o.PortName = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *Endpoint) GetInstances() []string {
	if o == nil || o.Instances == nil {
		var ret []string
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetInstancesOk() (*[]string, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return &o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *Endpoint) HasInstances() bool {
	return o != nil && o.Instances != nil
}

// SetInstances gets a reference to the given []string and assigns it to the Instances field.
func (o *Endpoint) SetInstances(v []string) {
	o.Instances = v
}

// GetMutable returns the Mutable field value.
func (o *Endpoint) GetMutable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Mutable
}

// GetMutableOk returns a tuple with the Mutable field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetMutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mutable, true
}

// SetMutable sets field value.
func (o *Endpoint) SetMutable(v bool) {
	o.Mutable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["component"] = o.Component
	toSerialize["hosts"] = o.Hosts
	toSerialize["port"] = o.Port
	toSerialize["type"] = o.Type
	toSerialize["networkType"] = o.NetworkType
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["portName"] = o.PortName
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	toSerialize["mutable"] = o.Mutable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Endpoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title       *string              `json:"title"`
		Component   *string              `json:"component"`
		Hosts       *[]string            `json:"hosts"`
		Port        *int32               `json:"port"`
		Type        *EndpointType        `json:"type"`
		NetworkType *EndpointNetworkType `json:"networkType"`
		ServiceName *string              `json:"serviceName"`
		PortName    *string              `json:"portName"`
		Instances   []string             `json:"instances,omitempty"`
		Mutable     *bool                `json:"mutable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Hosts == nil {
		return fmt.Errorf("required field hosts missing")
	}
	if all.Port == nil {
		return fmt.Errorf("required field port missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.NetworkType == nil {
		return fmt.Errorf("required field networkType missing")
	}
	if all.ServiceName == nil {
		return fmt.Errorf("required field serviceName missing")
	}
	if all.PortName == nil {
		return fmt.Errorf("required field portName missing")
	}
	if all.Mutable == nil {
		return fmt.Errorf("required field mutable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "component", "hosts", "port", "type", "networkType", "serviceName", "portName", "instances", "mutable"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Title = *all.Title
	o.Component = *all.Component
	o.Hosts = *all.Hosts
	o.Port = *all.Port
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.NetworkType.IsValid() {
		hasInvalidField = true
	} else {
		o.NetworkType = *all.NetworkType
	}
	o.ServiceName = *all.ServiceName
	o.PortName = *all.PortName
	o.Instances = all.Instances
	o.Mutable = *all.Mutable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
