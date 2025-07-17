// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeServiceRefHelmValuePath The path to be used in values. Separated with commas. ClusterCreate API will use these path to override values in the cluster chart.
type ModeServiceRefHelmValuePath struct {
	// the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.
	Namespace string `json:"namespace"`
	// the name of the referenced Cluster
	Cluster string `json:"cluster"`
	// see serviceSelectors
	Component *string `json:"component,omitempty"`
	// see serviceSelectors
	Service *string `json:"service,omitempty"`
	// see serviceSelectors
	Port *string `json:"port,omitempty"`
	// the name of the referenced serviceDescriptor
	ServiceDescriptor string `json:"serviceDescriptor"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeServiceRefHelmValuePath instantiates a new ModeServiceRefHelmValuePath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeServiceRefHelmValuePath(namespace string, cluster string, serviceDescriptor string) *ModeServiceRefHelmValuePath {
	this := ModeServiceRefHelmValuePath{}
	this.Namespace = namespace
	this.Cluster = cluster
	this.ServiceDescriptor = serviceDescriptor
	return &this
}

// NewModeServiceRefHelmValuePathWithDefaults instantiates a new ModeServiceRefHelmValuePath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeServiceRefHelmValuePathWithDefaults() *ModeServiceRefHelmValuePath {
	this := ModeServiceRefHelmValuePath{}
	return &this
}

// GetNamespace returns the Namespace field value.
func (o *ModeServiceRefHelmValuePath) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *ModeServiceRefHelmValuePath) SetNamespace(v string) {
	o.Namespace = v
}

// GetCluster returns the Cluster field value.
func (o *ModeServiceRefHelmValuePath) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *ModeServiceRefHelmValuePath) SetCluster(v string) {
	o.Cluster = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ModeServiceRefHelmValuePath) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ModeServiceRefHelmValuePath) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ModeServiceRefHelmValuePath) SetComponent(v string) {
	o.Component = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ModeServiceRefHelmValuePath) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ModeServiceRefHelmValuePath) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ModeServiceRefHelmValuePath) SetService(v string) {
	o.Service = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModeServiceRefHelmValuePath) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModeServiceRefHelmValuePath) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ModeServiceRefHelmValuePath) SetPort(v string) {
	o.Port = &v
}

// GetServiceDescriptor returns the ServiceDescriptor field value.
func (o *ModeServiceRefHelmValuePath) GetServiceDescriptor() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceDescriptor
}

// GetServiceDescriptorOk returns a tuple with the ServiceDescriptor field value
// and a boolean to check if the value has been set.
func (o *ModeServiceRefHelmValuePath) GetServiceDescriptorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDescriptor, true
}

// SetServiceDescriptor sets field value.
func (o *ModeServiceRefHelmValuePath) SetServiceDescriptor(v string) {
	o.ServiceDescriptor = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeServiceRefHelmValuePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["namespace"] = o.Namespace
	toSerialize["cluster"] = o.Cluster
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	toSerialize["serviceDescriptor"] = o.ServiceDescriptor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeServiceRefHelmValuePath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace         *string `json:"namespace"`
		Cluster           *string `json:"cluster"`
		Component         *string `json:"component,omitempty"`
		Service           *string `json:"service,omitempty"`
		Port              *string `json:"port,omitempty"`
		ServiceDescriptor *string `json:"serviceDescriptor"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	if all.ServiceDescriptor == nil {
		return fmt.Errorf("required field serviceDescriptor missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"namespace", "cluster", "component", "service", "port", "serviceDescriptor"})
	} else {
		return err
	}
	o.Namespace = *all.Namespace
	o.Cluster = *all.Cluster
	o.Component = all.Component
	o.Service = all.Service
	o.Port = all.Port
	o.ServiceDescriptor = *all.ServiceDescriptor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
